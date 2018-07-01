package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"os/exec"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Command string `json:"command"`
}

type Response struct {
	Command string `json:"response"`
}

func main() {

	http.HandleFunc("/", handleConnections)

	//go handleMessages()

	log.Println("python-module started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	// upgrade the GET request to a websocket; else, exit
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// close the connection when the function returns
	defer ws.Close()

	// register the client
	clients[ws] = true

	// loop over each message gotten
	for {
		var msg Message

		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		cmd := exec.Command("python", "-c", msg.Command)
		cmdOut, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		ws.WriteJSON(Response{string(cmdOut)})

		//broadcast <- msg
	}
}
