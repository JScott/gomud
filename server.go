package main

import (
	"fmt"
	"time"

	zmq "github.com/alecthomas/gozmq"
	"github.com/fiorix/go-redis/redis"
)

const noFlags = 0

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REP)
	defer context.Close()
	defer socket.Close()
	socket.Bind("tcp://*:5555")

	storage := redis.New("localhost:6379")

	for {
		msg, _ := socket.Recv(noFlags)
		println("Received ", string(msg))

		time.Sleep(time.Second)

		reply := fmt.Sprintf("World")
		socket.Send([]byte(reply), noFlags)
	}
}
