package main

import (
	"fmt"

	"github.com/chrisbendel/go_ws/utils"
	"golang.org/x/net/websocket"
)

var clients []Client
var currentWord = utils.GetRandomWord()
var wsHandler = websocket.Handler(onWsConnect)

func onWsConnect(ws *websocket.Conn) {
	defer ws.Close()
	client := NewClient(ws)
	client.name = fmt.Sprintf("%d", len(clients)+1)
	clients = addClientAndGreet(clients, client)
	client.listen()
}

func broadcast(msg *Message) {
	fmt.Printf("Broadcasting %+v\n", msg)
	for _, c := range clients {
		c.ch <- msg
	}
}

func addClientAndGreet(list []Client, client Client) []Client {
	hiddenWord := utils.ReplaceLetters(currentWord, client.guessed)
	clients = append(list, client)
	websocket.JSON.Send(client.connection, Message{fmt.Sprintf("Hey welcome to the coolest hangman game ever. You are currently player %s", client.name), fmt.Sprintf("Your current word to guess is %s", hiddenWord)})
	return clients
}
