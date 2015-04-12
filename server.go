package main

import (
  "fmt"
  "time"
  "log"
  //"errors"

  "gomud/connections"
  //"encoding/json"
  //"github.com/fiorix/go-redis/redis"
  zmq "github.com/pebbe/zmq4"
)

const noFlags = 0

func main() {
  server, err := connections.CreateServer()
  defer server.Close()
  if err == nil {
    waitForPlayersOn(server)
  } else {
    log.Fatal(err)
  }
}

func waitForPlayersOn(server *zmq.Socket) {
  //storage := redis.New("localhost:6379")
  for {
    //  Wait for next request from client
    msg, _ := server.Recv(noFlags)
    fmt.Println("Received ", msg)

    //  Do some 'work'
    time.Sleep(time.Second)

    //  Send reply back to client
    reply := "World"
    server.Send(reply, noFlags)
  }
}
