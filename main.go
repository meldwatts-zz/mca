package main

import (
	"log"
	"net/http"
)

const (
	connHost = "localhost" // Dirección de acceso
	connPort = "8080"      // Puerto de escucha
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/registrar.html")
}

func registro(w http.ResponseWriter, r *http.Request) {
	ingreso := r.PostFormValue("ingreso")
	ingresar := r.PostFormValue("ingresar")
	egresar := r.PostFormValue("egresar")
	if ingresar == "ingresar" {
		log.Println(ingreso)
		log.Println(ingresar)
		http.ServeFile(w, r, "templates/ingreso.html")
	}
	if egresar == "egresar" {
		log.Println(ingreso)
		log.Println(egresar)
	}

}

func main() {
	//Log del servidor
	log.Println("Ejecutando Servidor en: " + connPort)
	//Manejadores
	http.HandleFunc("/registro", registro)
	http.HandleFunc("/", index)
	//Manejador de los archivos estaticos
	http.Handle("/public/static/", http.StripPrefix("/public/static/", http.FileServer(http.Dir("public/static"))))
	//Ejecución del Servidor y log
	log.Fatal(http.ListenAndServe(connHost+":"+connPort, nil))
}
