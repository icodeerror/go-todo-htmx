package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todos struct {
	ID          int
	Description string
	Completed   bool
}

func main() {
	srv := http.NewServeMux()

	srv.HandleFunc("GET /{$}", homeHandler)
	srv.HandleFunc("GET /todo/{id}", getTodo)

	fmt.Println("Server is starting on http://localhost:8000")
	err := http.ListenAndServe(":8000", srv)
	if err != nil {
		log.Fatal(err)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// t, err := template.ParseFiles("./web/base.html", "./web/index.html")
	todos := []Todos{
		{
			ID:          1,
			Description: "Buy something now",
			Completed:   false,
		},
		{
			ID:          2,
			Description: "Buy something later",
			Completed:   false,
		},
		{
			ID:          3,
			Description: "Buy something please",
			Completed:   true,
		},
	}

	t, err := template.ParseGlob("./web/*.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "base", todos)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {

	urlVal := r.PathValue("id")

	w.Write([]byte(urlVal))
}

func markAsCompleted(w http.ResponseWriter, r *http.Request) {
	id := r.Url
}
