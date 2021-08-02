package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gavr-games/reborn-mmorpg/pkg/game"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	engine := game.NewEngine()
	go engine.Run()
	go engine.RunWorld()
	http.HandleFunc("/engine/ws", func(w http.ResponseWriter, r *http.Request) {
		game.ServeWs(engine, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
