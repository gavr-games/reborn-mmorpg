package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gavr-games/reborn-mmorpg/pkg/chat"
)

var addr = flag.String("addr", ":8081", "http service address")

func main() {
	flag.Parse()
	hub := chat.NewHub()
	go hub.Run()
	http.HandleFunc("/chat/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
