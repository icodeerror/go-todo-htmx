package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles("./web/base.layout.html", fmt.Sprintf("./web/%s", page))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func renderHTMXTemplate(w http.ResponseWriter, partTemplate string, data any) {
	tmpl, err := template.ParseFiles(fmt.Sprintf("./web/%s", partTemplate))
	// tmpl, err := template.ParseGlob("./web/*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, partTemplate, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTodoID(r *http.Request) (int, error) {
	// Get id from url
	pathVal := r.PathValue("id")

	// Convert ascii to integer
	id, err := strconv.Atoi(pathVal)
	if err != nil || id < 1 {
		return 0, err
	}

	return id, nil
}

// func (app *app) renderTemplate(w http.ResponseWriter, page string, data any) {
// 	buf := new(bytes.Buffer)

// 	t, ok := app.templateCache[page]
// 	if !ok {
// 		fmt.Printf("template %s doesn't exist", page)
// 		return
// 	}

// 	err := t.ExecuteTemplate(buf, "base", data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	buf.WriteTo(w)

// }

// func newTemplateCache() (map[string]*template.Template, error) {
// 	cache := map[string]*template.Template{}

// 	pages, err := filepath.Glob("./web/*.page.html")
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)

// 		files := []string{
// 			"./web/base.layout.html",
// 			// "./web/todo.part.html",
// 			page,
// 		}
// 		ts, err := template.ParseFiles(files...)
// 		if err != nil {
// 			return nil, err
// 		}

// 		cache[name] = ts

// 	}

// 	fmt.Println("Cache", cache)
// 	return cache, nil

// }
