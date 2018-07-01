package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Command string `json:"command"`
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

		cmd = exec.Command("python", "-m", msg.Command)
		cmdOut = cmd.Output()
		if err != nil {
			panic(err)
		}
		ws.WriteJSON(msg(string(cmdOut)))

		//broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// package main

// import "fmt"
// import "io/ioutil"
// import "os/exec"

// func main() {

// 	dateCmd := exec.Command("date")

// 	dateOut, err := dateCmd.Output()
// 	if err != nul {
// 		panic(err)
// 	}
// 	fmt.Println("> date")
// 	fmt.Println(string(dateOut))

// 	grepCmd := exec.Command("grep", "hello")

// 	grepIn, _ := grepCmd.StdinPipe()
// 	grepOut, _ := grepCmd.StdoutPipe()
// 	grepCmd.Start()
// 	grepIn.Write([]byte("hello grep\ngoodbye grep"))
// 	grepIn.Close()
// 	grepBytes, _ := ioutil.ReadAll(grepOut)
// 	grepCmd.Wait()

// 	fmt.Println("> grep hello")
// 	fmt.Println(string(grepBytes))

// 	lsCmd = exec.Command("bash", "-c", "ls -a -l -h")
// 	lsOut, err := lsCmd.Output()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("> ls -a -l -h")
// 	fmt.Println(string(lsOut))
// }
	
	
