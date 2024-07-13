const { ReadItemRequest, MessageRequest, APIServiceClient } = require('./web_grpc_web_pb.js');

var google_protobuf_empty_pb  = require('google-protobuf/google/protobuf/empty_pb.js');

var client;
var token;
window.NewClient = function(proxyAddr_, token_){
  client = new APIServiceClient(proxyAddr_, null, null);
  token = token_;
  return client;
}

window.EmptyRequest = function(){
  const request = new google_protobuf_empty_pb.Empty();
  return request;
}

window.ReadItem = function(method, id = -1) {
  return new Promise((resolve, reject) => {
    var request = new ReadItemRequest();
    request.setMethod(method);
    request.setId(id);

    var metadata = {'Authorization': 'Bearer '+ token};
    client.readItem(request, metadata, (err, response) => {
      if (err) {
        reject(err);
      } 

      const status = response.getStatus();
      const data = response.getData();
      resolve({
        status: status,
        data: data,
      });
    })
  });
};

window.DeleteItem = function(method, id = -1) {
  return new Promise((resolve, reject) => {
    var request = new ReadItemRequest();
    request.setMethod(method);
    request.setId(id);

    var metadata = {'Authorization': 'Bearer '+ token};
    client.deleteItem(request, metadata, (err, response) => {
      if (err) {
        reject(err);
      } 

      const status = response.getStatus();
      const data = response.getData();
      resolve({
        status: status,
        data: data,
      });
    });
  });
};


window.ShowClock = function(){
  const request = EmptyRequest();

  // clock流
  const stream = client.showClock(request, {});
  const clock = document.getElementById('clock');

  // 監聽流的數據事件
  stream.on('data', function(response) {
    clock.innerHTML = response.getData();
  });

  stream.on('error', function(error){
    //console.log("流錯誤:", error);
    clock.innerHTML = '時鐘加載失敗';
  })

  // 監聽流的结束事件
  stream.on('end', function() {
    console.log('時鐘流结束');
  });
}


window.SendMessage = function(action, userId, message) {
  return new Promise((resolve, reject) => {
    var request = new MessageRequest();
    request.setAction(action);
    request.setUserid(userId);
    request.setMessage(message);

    var metadata = {'Authorization': 'Bearer '+ token};
    client.sendMessage(request, metadata, (err, response) => {
      if (err) {
        reject(err);
      } 

      const status = response.getStatus();
      const data = response.getData();
      resolve({
        status: status,
        data: data,
      });
    });
  });
};
