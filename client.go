package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	zmq "github.com/pebbe/zmq4"
)

const noFlags = 0

func main() {
	//  Socket to talk to server
	fmt.Println("Connecting...")
	client := createClient()
	if client != nil {
		readCommandsInto(client)
		client.Close()
	}
}

func createClient() *zmq.Socket {
	client, err := zmq.NewSocket(zmq.REQ)
	if err == nil {
		client.Connect("tcp://localhost:5555")
		return client
	} else {
		// TODO: go style error standards?
		log.Fatal(err)
		return nil
	}
}

func readCommandsInto(client *zmq.Socket) {
	stream := bufio.NewScanner(os.Stdin)
	for stream.Scan() {
		sendMessage(stream.Text(), client)
	}
	err := stream.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func sendMessage(message string, client *zmq.Socket) {
	fmt.Println("Sending ", message)
	client.Send(message, noFlags)
	reply, err := client.Recv(0)
	if err == nil {
		fmt.Println("Received ", reply)
	} else {
		log.Fatal(err)
	}
}
