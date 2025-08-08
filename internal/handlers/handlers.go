package handlers

import (
    "net/http"
    "github.com/a-h/templ"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    templ.Handler(Home()).ServeHTTP(w, r)
}

func PartialHandler(w http.ResponseWriter, r *http.Request) {
    templ.Handler(Partial()).ServeHTTP(w, r)
}
