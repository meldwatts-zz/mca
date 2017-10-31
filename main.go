package main

import (
	"log"
	"net/http"
)

func registro(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func main() {
	port := ":8080"
	log.Println("Ejecutando Servidor en: " + port)
	http.HandleFunc("/", registro)
	log.Fatal(http.ListenAndServe(port, nil))
}
