package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Parsear el archivo index.html desde la carpeta templates
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error al cargar el template:", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	// Ejecutar el template y escribirlo en la respuesta HTTP
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error al ejecutar el template:", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
	}
}

func main() {
	// Servir archivos estáticos de la carpeta 'images'
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("images"))))

	// Ruta para la página principal
	http.HandleFunc("/", handler)

	log.Println("Servidor escuchando en puerto 8080...")
	http.ListenAndServe(":8080", nil)
}
