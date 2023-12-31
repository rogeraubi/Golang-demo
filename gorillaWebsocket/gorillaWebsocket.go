package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
func main ()  {
	   http.HandleFunc("/echo", func (w http.ResponseWriter, r *http.Request)  {
		conn, _ := upgrader.Upgrade(w,r,nil)      
		for {
			 msgType, msg , err := conn.ReadMessage()
			 if err != nil {
				return 
			 }
             fmt.Printf("%s -> %s sent: %s\n", conn.RemoteAddr(),string(msg),string("1111"))

			 if err = conn.WriteMessage(msgType, msg); err != nil {
				return 
			 }
		}
		// fmt.Printf("%s sent: %s\n" , conn.Remote)
	})

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		http.ServeFile(w,r,"websocket.html")
	})
	http.ListenAndServe(":8080",nil)
}
/**
$ go run websockets.go
[127.0.0.1]:53403 sent: Hello Go Web Examples, you're doing great!

**/