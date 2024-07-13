package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	*websocket.Conn
}

type WsPayload struct {
	Action      string              `json:"action"`
	Message     string              `json:"message"`
	UserName    string              `json:"username"`
	MessageType string              `json:"message_type"`
	UserID      int64               `json:"user_id"`
	Conn        WebSocketConnection `json:"-"`
}

type WsJsonResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	UserID  int64  `json:"user_id"`
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var clients = make(map[WebSocketConnection]string)

var wsChan = make(chan WsPayload)

func (app *application) WsEndPoint(w http.ResponseWriter, r *http.Request) {
	wupgrade := w
	if u, ok := w.(interface{ Unwrap() http.ResponseWriter }); ok {
		wupgrade = u.Unwrap()
	}

	conn, err := upgradeConnection.Upgrade(wupgrade, r, nil)
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	defer conn.Close()

	app.infoLog.Println("WsEndPoint connected")

	var response WsJsonResponse
	response.Message = "Connected to server"

	err = conn.WriteJSON(response)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	ws := WebSocketConnection{Conn: conn}
	clients[ws] = ""

	app.ListenForWS(&ws)
}

func (app *application) ListenForWS(conn *WebSocketConnection) {
	defer conn.Close()

	defer func() {
		if r := recover(); r != nil {
			app.errorLog.Println("ERORR:", fmt.Sprintf("%v", r))
			// Close the connection or perform any other cleanup
			// conn.Close()
		}
	}()

	var payload WsPayload
	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			fmt.Println("read:", payload)
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func (app *application) ListenToWsChannel() {
	var response WsJsonResponse
	for {
		e := <-wsChan
		switch e.Action {
		case "deleteUser":
			response.Action = "deleteUser"
			response.Message = "Your account has been deleted"
			response.UserID = e.UserID
			app.broadcastToAll(response)
		case "logout":
			response.Action = "logout"
			response.Message = ""
			response.UserID = e.UserID
			app.broadcastToAll(response)
		default:
		}
	}
}

func (app *application) broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		// broadcast to every connected client
		err := client.WriteJSON(response)
		if err != nil {
			app.errorLog.Printf("Websocket err on %s: %s", response.Action, err)
			_ = client.Close()
			delete(clients, client)
		} else {
			app.infoLog.Println("broadcast:", response)
		}
	}
}

func (app *application) WsNtpTime(w http.ResponseWriter, r *http.Request) {
	wupgrade := w
	if u, ok := w.(interface{ Unwrap() http.ResponseWriter }); ok {
		wupgrade = u.Unwrap()
	}

	conn, err := upgradeConnection.Upgrade(wupgrade, r, nil)
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	defer conn.Close()

	app.infoLog.Println("WsNtpTime connected")

	// 每秒向 WebSocket 客户端发送当前时间
	for {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		err := conn.WriteMessage(websocket.TextMessage, []byte(currentTime))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
