package main

import (
	"fmt"
	"time"

	"github.com/fiorix/go-redis/redis"
	zmq "github.com/pebbe/zmq4"
)

const noFlags = 0

func main() {
	server, _ := zmq.NewSocket(zmq.REP)
	defer server.Close()
	server.Bind("tcp://*:5555")

	storage := redis.New("localhost:6379")

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
