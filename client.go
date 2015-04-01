package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	zmq "github.com/pebbe/zmq4"
	"gomud/connections"
)

const usageErr = "client.go <user> <pass>"
const authErr = "Credentials failed to log in"

func main() {
	//  Socket to talk to server
	fmt.Println("Connecting...")
	client, err := connections.CreateClient()
	defer client.Close()
	if err == nil {
		loginAndPlay(client)
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

func requestLoginWith(client *zmq.Socket) (string, error) {
	if len(os.Args) < 2 {
		log.Fatal(usageErr)
	}
	username := os.Args[1]
	//password := os.Args[2]
  return connections.RequestLogin(username, "", client)
}

func readCommandsInto(client *zmq.Socket, authToken string) {
	stream := bufio.NewScanner(os.Stdin)
	for stream.Scan() {
		handlePlayerInput(stream.Text(), authToken, client)
	}
	err := stream.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func handlePlayerInput(input string, authToken string, client *zmq.Socket) {
	split := strings.SplitN(input, " ", 2)
	command := split[0]
	body := ""
	if len(split) > 1 {
		body = split[1]
	}
	reply, err := connections.RequestAction(command, body, authToken, client)
	if err == nil {
		fmt.Println("Received ", reply)
	} else {
		log.Fatal(err)
	}
}
