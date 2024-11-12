package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Ruta para servir archivos estáticos (como imágenes)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates"))))

	// Ruta para servir la página principal con el template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	log.Println("Servidor en http://localhost:3002")
	http.ListenAndServe(":3002", nil)
}
