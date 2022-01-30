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
}

func homePage(res http.ResponseWriter, req *http.Request) {

}

func Reader(conn *websocket.Conn) {
	//always do the loop and read the message when the endpoint is getting hit
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %s", err.Error())
			return
		}

		log.Printf("message: %s", string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Error writing message: %s", err.Error())
			return
		}
	}

}

func wsFunction(res http.ResponseWriter, req *http.Request) {

	log.Println("hitting ws enpoint")
	// allow anything, from any website origins, to connect to websocket connection
	// dont do this on real projects
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	//upgrading connection from http/ip connection to 101 websocket connection
	wsConn, err := upgrader.Upgrade(res, req, nil)

	if err != nil {
		log.Printf("error upgrading connection: %s", err.Error())
		return
	}
	fmt.Println("success connecting clients....")

	Reader(wsConn)

	defer wsConn.Close()
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsFunction)
}

func main() {

	setupRoutes()
	if err := http.ListenAndServe(":8100", nil); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}

}
