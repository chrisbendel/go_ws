package main

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/chrisbendel/go_ws/letters"
	"github.com/chrisbendel/go_ws/words"
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
			fmt.Printf("Received: %+v\n", msg.Body)
			reg, err := regexp.Compile("[^a-zA-Z0-9]+")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("hello1")
			processedString := reg.ReplaceAllString(msg.Body, "")
			guessedRune := rune('a')

			if processedString != "" {
				guessedRune = rune(processedString[0])
			}

			if !strings.ContainsRune(string(c.guessed), guessedRune) {
				c.guessed = append(c.guessed, guessedRune)
			}

			userGuess := letters.ReplaceLetters(currentWord, c.guessed)

			if err == io.EOF {
				c.close <- true
			} else if err != nil {
				// c.server.Err(err)
			} else {
				fmt.Println(&msg)

				if letters.IsCorrect(currentWord, userGuess) {
					currentWord = words.GetRandomWord()
					c.guessed = make([]rune, 0)
					m := Message{
						Author: "Server",
						Body:   fmt.Sprintf("Player %s won the game. Start typing to play the next round. Your hint is %s", c.name, letters.ReplaceLetters(currentWord, c.guessed)),
					}
					broadcast(&m)
				} else {
					m := Message{
						Author: c.name,
						Body:   userGuess,
					}

					c.ch <- &m
				}
			}
		}
	}
}
