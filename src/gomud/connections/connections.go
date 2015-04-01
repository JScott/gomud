package connections

import (
  "encoding/json"
  "fmt"
  zmq "github.com/pebbe/zmq4"
)

const noFlags = 0
const serverAddress = "tcp://localhost:5555"
const serverBinding = "tcp://*:5555"

func CreateClient() (*zmq.Socket, error) {
  client, err := zmq.NewSocket(zmq.REQ)
  if err == nil {
    client.Connect(serverAddress)
  }
  return client, err
}

func CreateServer() (*zmq.Socket, error) {
  server, err := zmq.NewSocket(zmq.REP)
  if err == nil {
    server.Bind(serverBinding)
  }
  return server, err
}

func RequestLogin(username string, password string, client *zmq.Socket) (string, error) {
  fmt.Println("Logging in as", username)
  return sendCommand("login", username+":"+password, "", client)
}

func RequestAction(action string, body string, authToken string, client *zmq.Socket) (string, error) {
  fmt.Println("Sending ", action)
  return sendCommand(action, body, authToken, client)
}

func sendCommand(action string, body string, authToken string, client *zmq.Socket) (string, error) {
  message, err := serializeCommand(action, body, authToken)
  if err == nil {
    client.Send(message, noFlags)
    return client.Recv(noFlags)
  } else {
    return "", err
  }
}

func serializeCommand(action string, body string, authToken string) (string, error) {
  messageMap := map[string]string{
    "action": action,
    "body": body,
    "authToken": authToken,
  }
  bytes, err := json.Marshal(messageMap)
  return string(bytes), err
}
