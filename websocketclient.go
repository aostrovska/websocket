
package main

import (
	"log"
	"time"
	"github.com/gorilla/websocket"
)
 

func read(c *websocket.Conn){
	for{
		_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
		log.Printf("user: %s", message)
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
	for {
		err := c.WriteMessage(websocket.TextMessage, []byte("Hi, I'm client 1"))
		if err != nil {
				log.Println("read:", err)
				return
			}
		err = c.WriteMessage(websocket.TextMessage, []byte("Hi, I'm client 2"))
		if err != nil {
				log.Println("read:", err)
				return
			}
	}
	/*	_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		time.Sleep(1*time.Second)
	}*/

	
}
