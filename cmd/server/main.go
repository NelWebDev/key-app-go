package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
        mensaje := struct {
            Mensaje string `json:"mensaje"`
        }{
            Mensaje: "¡Hola, mundo!",
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(mensaje)
    })

    log.Fatal(http.ListenAndServe(":8080", router))
}