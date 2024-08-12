package main

import (
	"log"
	"net/http"

	"github.com/icodeerror/go-todo-htmx/internal/models"
)

func (app *app) homeHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todoModel.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	renderTemplate(w, "index.page.html", todos)
}

func (app *app) getEditTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getTodoID(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	todo, err := app.todoModel.GetID(id)
	if err != nil {
		// Return not found if row doesn't exist
		http.NotFound(w, r)
		return
	}

	renderHTMXTemplate(w, "update_form.part.html", todo)
}

func (app *app) markCompleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getTodoID(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Get single todo
	todo, err := app.todoModel.GetID(id)
	if err != nil {
		log.Fatal(err)
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPatch {
		renderTemplate(w, "index.page.html", nil)
		return
	}

	// Which todo to be mark completed
	updateTodo := models.Todos{
		ID:          id,
		Description: todo.Description,
		Completed:   !todo.Completed,
	}

	// Mark complete selected todo
	err = app.todoModel.MarkComplete(updateTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderHTMXTemplate(w, "todo_single.part.html", updateTodo)

}

func (app *app) updateDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getTodoID(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Get single todo
	todo, err := app.todoModel.GetID(id)
	if err != nil {
		log.Fatal(err)
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPatch {
		renderTemplate(w, "index.page.html", nil)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Limit form to 1024 bytes
	r.Body = http.MaxBytesReader(w, r.Body, 1024)

	description := r.PostForm.Get("description")
	if description == "" {
		description = todo.Description
	}
	// Pass todo ID and update
	updateTodo := models.Todos{
		ID:          todo.ID,
		Description: description,
		Completed:   todo.Completed,
	}

	err = app.todoModel.UpdateDescription(updateTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderHTMXTemplate(w, "update_done.part.html", updateTodo)

}

func (app *app) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get initial todos
	todos, err := app.todoModel.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get input with name "description"
	description := r.PostForm.Get("description")

	// If empty, render todo list with initial todos
	if description == "" {
		renderHTMXTemplate(w, "todo_list.part.html", todos)
		return
	}

	// Assign new todo input to variable with default "Completed" false
	newTodo := models.Todos{
		Description: description,
		Completed:   false,
	}

	// Insert new record to the table
	err = app.todoModel.Insert(newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch again all todos, so it will display newly added todo
	todos, err = app.todoModel.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Render todo list
	renderHTMXTemplate(w, "todo_list.part.html", todos)
}

func (app *app) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getTodoID(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = app.todoModel.Delete(models.Todos{ID: id})
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Fetch all todos
	todos, err := app.todoModel.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderHTMXTemplate(w, "todo_list.part.html", todos)
}

func (app *app) cancelUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := getTodoID(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Get single todo
	todo, err := app.todoModel.GetID(id)
	if err != nil {
		log.Fatal(err)
		http.NotFound(w, r)
		return
	}
	renderHTMXTemplate(w, "update_done.part.html", todo)
}
