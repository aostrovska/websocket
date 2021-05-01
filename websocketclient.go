
package main

import (
	"log"
	"github.com/gorilla/websocket"
	"time"
	"fmt"
)
 

func read(c *websocket.Conn){
	for{
		_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
		log.Printf("user: %s", message)
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
		var mes string
		fmt.Scan(&mes)
		err := c.WriteMessage(websocket.TextMessage, []byte(mes))
		if err != nil {
				log.Println("read:", err)
				return
			}
		time.Sleep(1*time.Second)
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
