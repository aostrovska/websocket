
package main

import (
	"log"
	"time"
	"github.com/gorilla/websocket"
)

var broadcast = make(chan int) 

func send(mes []byte, c *websocket.Conn){
	err := c.WriteMessage(websocket.TextMessage, mes)
	if err != nil {
			log.Println("read:", err)
			return
	}
	log.Printf("client 1: %s", mes)
	broadcast <-1
	time.Sleep(1*time.Second)
}

func read(c *websocket.Conn){
	for{
		_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
		log.Printf("client 2: %s", message)
		time.Sleep(1*time.Second)
	}
}

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/socket", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go read(c)
	go send([]byte("Hi, I'm client 1"), c)
	go send ([]byte("Hi, I'm client 2"), c)
	/*for {
		err := c.WriteMessage(websocket.TextMessage, []byte("Hello server"))
		
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		time.Sleep(1*time.Second)
	}*/

	
}
