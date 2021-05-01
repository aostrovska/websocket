package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"log"
	"github.com/gorilla/websocket"
)
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte) 
var upgrader = websocket.Upgrader{}

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }
		
		log.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	} else if req.Method == "OPTIONS" {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(405)
	}
	
}

func Socket(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	clients[conn] = true
	
	for {
		_, mes, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		broadcast <- mes
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
				return
			}
		}
	}
}
func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/socket", Socket)
	
	go handleMessages()
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}

