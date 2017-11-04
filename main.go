package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/registrar.html")
}

func registro(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte(r.URL.Path))
	//r.ParseForm()
	ingreso := r.PostFormValue("ingreso")
	ingresar := r.PostFormValue("ingresar")
	egresar := r.PostFormValue("egresar")
	//var message string
	if ingresar == "ingresar" {
		log.Println(ingreso)
		log.Println(ingresar)
		//	message = "Ingresar Visitante"
		http.ServeFile(w, r, "templates/ingreso.html")
	}
	if egresar == "egresar" {
		log.Println(ingreso)
		log.Println(egresar)
		//	message = "Salida de Visitante"
	}

	//w.Write([]byte(message))
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
