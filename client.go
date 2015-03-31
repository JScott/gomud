package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"errors"
	"encoding/json"

	zmq "github.com/pebbe/zmq4"
)

const noFlags = 0
const usageErr = "client.go <user> <pass>"
const authErr = "Credentials failed to log in"
const serverAddress = "tcp://localhost:5555"

func main() {
	//  Socket to talk to server
	fmt.Println("Connecting...")
	client, err := createClient()
	if err == nil {
		loginAndPlay(client)
		client.Close()
	} else {
		log.Fatal(err)
	}
}

func loginAndPlay(client *zmq.Socket) {
	token, err := requestLoginWith(client)
	if err == nil {
		readCommandsInto(client, token)
	} else {
		log.Fatal(err)
	}
}

func createClient() (*zmq.Socket, error) {
	client, err := zmq.NewSocket(zmq.REQ)
	if err == nil {
		client.Connect(serverAddress)
		return client, nil
	} else {
		return nil, err
	}
}

func requestLoginWith(client *zmq.Socket) (string, error) {
	if len(os.Args) < 2 {
		log.Fatal(usageErr)
	}
	username := os.Args[1]

  reply, err := sendMessage(username, client)
  if err == nil {
  	if reply == "denied" {
  		return reply, errors.New(authErr)
  	} else {
  		return reply, nil
  	}
  } else {
  	return reply, err
  }
}

func readCommandsInto(client *zmq.Socket, authToken string) {
	stream := bufio.NewScanner(os.Stdin)
	for stream.Scan() {
		message, err := generateMessage(authToken, stream.Text())
		if err == nil {

			// TODO: new method
			reply, err := sendMessage(message, client)
			if err == nil {
				fmt.Println("Received ", reply)
			} else {
				log.Fatal(err)
			}

		} else {
			log.Fatal(err)
		}
	}
	err := stream.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func generateMessage(authToken string, message string) (string, error) {
	messageMap := map[string]string{
		"authToken": authToken,
		"message": message,
	}
	bytes, err := json.Marshal(messageMap)
	return string(bytes), err
}

func sendMessage(message string, client *zmq.Socket) (string, error) {
	fmt.Println("Sending ", message)
	client.Send(message, noFlags)
	return client.Recv(0)
}
