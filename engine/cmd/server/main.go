package main

import (
	"flag"
	"net/http"
	"github.com/gavr-games/reborn-mmorpg/pkg/game"
	//"os"
  //"runtime/pprof"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	// Create a CPU profile file
	/*f, err := os.Create("profile.prof")
	if err != nil {
			panic(err)
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
	}
	defer pprof.StopCPUProfile()*/

	flag.Parse()
	engine := game.NewEngine()
	http.HandleFunc("/engine/ws", func(w http.ResponseWriter, r *http.Request) {
		game.ServeWs(engine, w, r)
	})
	go http.ListenAndServe(*addr, nil)
	engine.Run()
}
