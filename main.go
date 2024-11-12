package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Hola Mundo en GO", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	log.Println("Servidor en ejecución en http://localhost:3002")
	log.Fatal(http.ListenAndServe(":3002", nil))
}
