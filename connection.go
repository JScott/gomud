const serverAddress = "tcp://localhost:5555"
const serverBinding = "tcp://*:5555"

import (
  zmq "github.com/pebbe/zmq4"
)

func createClient() (*zmq.Socket, error) {
  client, err := zmq.NewSocket(zmq.REQ)
  defer client.Close()
  if err == nil {
    client.Connect(serverAddress)
  }
  return client, err
}

func createServer() (*zmq.Socket, error) {
  server, err := zmq.NewSocket(zmq.REP)
  defer server.Close()
  if err = nil {
    server.Bind(serverBinding)
  }
  return server, err
}

func sendMessage(message string, client *zmq.Socket) (string, error) {
  fmt.Println("Sending ", message)
  client.Send(message, noFlags)
  return client.Recv(0)
}

func generateMessage(authToken string, message string) (string, error) {
  messageMap := map[string]string{
    "authToken": authToken,
    "message": message,
  }
  bytes, err := json.Marshal(messageMap)
  return string(bytes), err
}
