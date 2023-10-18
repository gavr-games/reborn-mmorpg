package main

import (
    "fmt"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./assets/maps"))
    http.Handle("/engine_api/maps/", http.StripPrefix("/engine_api/maps/", fs))

    fmt.Println("Engine API started at port 8082")
    http.ListenAndServe(":8082", nil)
}
