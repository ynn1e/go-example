package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Message struct {
	Text string `json:"text"`
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.Handle("/chat", websocket.Handler(chatHandler))
	go handleMessages()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("listen and serve error: %s", err)
	}
}

func chatHandler(ws *websocket.Conn) {
	defer ws.Close()

	clients[ws] = true

LOOP:
	for {
		var msg Message
		err := websocket.JSON.Receive(ws, &msg)
		switch {
		case err == nil:
			broadcast <- msg
		case errors.Is(err, io.EOF):
			continue
		default:
			log.Printf("websocket receive error: %s", err)
			delete(clients, ws)
			break LOOP
		}
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		data, err := json.Marshal(msg)
		if err != nil {
			log.Printf("json marshal error: %s", err)
			continue
		}

		for client := range clients {
			if err := websocket.Message.Send(client, string(data)); err != nil {
				log.Printf("websocket send error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
