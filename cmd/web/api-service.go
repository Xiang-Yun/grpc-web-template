package main

import (
	"context"
	"net"
	"time"

	api_proto "grpc-web-template/cmd/web/api_proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	api_proto.UnimplementedAPIServiceServer
	App *application
}

func (app *application) startGrpcApiService() {
	lis, err := net.Listen("tcp", app.config.apiService)
	if err != nil {
		app.errorLog.Println("Failed to listen noise addr:", err)
		panic(err)
	}
	s := grpc.NewServer()

	api_proto.RegisterAPIServiceServer(s, &server{
		App: app,
	})

	app.infoLog.Println("api service listening on:", app.config.apiService)
	if err := s.Serve(lis); err != nil {
		app.errorLog.Println("Failed to api serve:", err)
		panic(err)
	}
}

// ShowClock show real time clock
func (s *server) ShowClock(in *emptypb.Empty, stream api_proto.APIService_ShowClockServer) error {
	for {
		if err := stream.Send(&api_proto.MessageResponse{
			Data: time.Now().Format("2006-01-02 15:04:05"),
		}); err != nil {
			break
		}
		time.Sleep(time.Millisecond * 990)
	}
	return nil
}

// SendMessage chat with client
func (s *server) SendMessage(ctx context.Context, msg *api_proto.MessageRequest) (*api_proto.MessageResponse, error) {
	if err := s.App.authenticateMetaToken(ctx); err != nil {
		return &api_proto.MessageResponse{Status: false}, err
	}

	response := &api_proto.MessageResponse{
		Status: true,
	}

	switch msg.GetAction() {
	case "logout":
		wsChan <- WsPayload{
			Action:  msg.GetAction(),
			Message: msg.GetMessage(),
			UserID:  msg.GetUserId(),
		}
	default:
	}

	return response, nil
}
