package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/registrar.html")
}

func registro(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte(r.URL.Path))
	r.ParseForm()
	var message string
	if r.Method == "GET" {
		message = "No Permitido"
	} else {
		message = "all pending tasks POST"
		ingreso := r.PostFormValue("ingreso")
		ingresar := r.PostFormValue("ingresar")
		egresar := r.PostFormValue("egresar")
		egresar2 := r.PostFormValue("egresar2")
		fmt.Println(ingreso)
		fmt.Println(ingresar)
		fmt.Println(egresar)
		fmt.Println(egresar2)
	}
	w.Write([]byte(message))

}

func main() {
	// Puerto de escucha
	port := ":8080"
	//Log del servidor
	log.Println("Ejecutando Servidor en: " + port)
	//Manejadores
	http.HandleFunc("/registro", registro)
	http.HandleFunc("/", index)
	//Manejador de los archivos estaticos
	http.Handle("/public/static/", http.StripPrefix("/public/static/", http.FileServer(http.Dir("public/static"))))
	//Ejecución del Servidor y log
	log.Fatal(http.ListenAndServe(port, nil))
}
