package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/ws", handleWebsocket)
	log.Fatal(http.ListenAndServe(":6969", nil))
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return

	}
	defer conn.Close()

	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received message: %s\n", data)
		fmt.Println("messageType:", messageType)
		if err := conn.WriteMessage(messageType, data); err != nil {
			fmt.Println(err)
			return
		}

	}
}
