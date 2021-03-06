package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/chrisbendel/go_ws/utils"
	"golang.org/x/net/websocket"
)

type Client struct {
	connection *websocket.Conn
	name       string
	ch         chan *Message
	close      chan bool
	guessed    []rune
}

func NewClient(ws *websocket.Conn) Client {
	ch := make(chan *Message, 100)
	close := make(chan bool)

	return Client{
		connection: ws,
		ch:         ch,
		close:      close,
	}
}

func (c *Client) listen() {
	go c.listenToWrite()
	c.listenToRead()
}

func (c *Client) listenToWrite() {
	for {
		select {
		case msg := <-c.ch:
			log.Println("Send:", msg)
			websocket.JSON.Send(c.connection, msg)

		case <-c.close:
			c.close <- true
			return
		}
	}
}

func (c *Client) listenToRead() {
	log.Println("Listening read from client")
	for {
		select {
		case <-c.close:
			c.close <- true
			return

		default:
			var msg Message
			err := websocket.JSON.Receive(c.connection, &msg)
			guessedRune := utils.UserGuessToRune(msg.Body)

			if !strings.ContainsRune(string(c.guessed), guessedRune) {
				c.guessed = append(c.guessed, guessedRune)
			}

			userGuess := utils.ReplaceLetters(currentWord, c.guessed)

			if err == io.EOF {
				c.close <- true
			} else if err != nil {
				// c.server.Err(err)
			} else {
				fmt.Println(&msg)
				previousWord := currentWord
				if utils.IsCorrect(currentWord, userGuess) {
					//Generate a new random word for the next round
					currentWord = utils.GetRandomWord()
					//Clear the user's previous guesses server side
					c.guessed = make([]rune, 0)
					//Message to inform the user that the game is over and notify everyone who won this round
					m := Message{
						Author: fmt.Sprintf("Player %s correctly guessed %s and won the game!!!", c.name, previousWord),
						Body:   fmt.Sprintf("Start typing to play the next round. Your hint is %s", utils.ReplaceLetters(currentWord, c.guessed)),
					}

					broadcast(&m)
				} else {
					m := Message{
						Author: fmt.Sprintf("You are player %s.", c.name),
						Body:   fmt.Sprintf("Current guess: %s", userGuess),
					}

					c.ch <- &m
				}
			}
		}
	}
}
