package main

import (
	"log"
	"net/http"
	"websocket/pkg/websocket"
)

func serveWs(p *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("websocket endpoint")

	conn, err := websocket.Upgrader(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: p,
	}

	p.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {

	setupRoutes()
	log.Println("server running http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
