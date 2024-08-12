package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/icodeerror/go-todo-htmx/internal/config"
	"github.com/icodeerror/go-todo-htmx/internal/models"
)

type app struct {
	todoModel models.TodosModel
	// templateCache map[string]*template.Template
}

// type Todos struct {
// 	ID          int
// 	Description string
// 	Completed   bool
// }

func main() {
	srv := http.NewServeMux()

	pool, db, err := config.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	defer db.Release()

	// templateCache, err := newTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	app := app{
		todoModel: models.TodosModel{DB: db},
		// templateCache: templateCache,
	}

	srv.HandleFunc("GET /{$}", app.homeHandler)
	srv.HandleFunc("PATCH /todo/{id}/complete", restrictDirectAccessHTMX(app.markCompleteHandler))
	srv.HandleFunc("GET /todo/{id}", restrictDirectAccessHTMX(app.getEditTodoHandler))
	srv.HandleFunc("GET /todo/{id}/cancel", restrictDirectAccessHTMX(app.cancelUpdate))
	srv.HandleFunc("PATCH /todo/{id}/update", restrictDirectAccessHTMX(app.updateDescriptionHandler))
	srv.HandleFunc("POST /todo", restrictDirectAccessHTMX(app.addTodoHandler))
	srv.HandleFunc("DELETE /todo/{id}/delete", restrictDirectAccessHTMX(app.deleteTodoHandler))

	fmt.Println("Server is starting on http://localhost:8000")
	err = http.ListenAndServe(":8000", bypassLocalTunnel(srv))
	if err != nil {
		log.Fatal(err)
	}
}
