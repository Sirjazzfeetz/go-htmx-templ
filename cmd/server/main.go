package main

import (
    "log"
    "net/http"
    "github.com/yourusername/go-htmx-templ/internal/handlers"
)

func main() {
    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/partial", handlers.PartialHandler)
    log.Println("Listening on http://localhost:3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
