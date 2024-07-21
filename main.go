package main

import (
	"html/template"
	"net/http"
	"sync"
)

// PageData holds the data to be rendered in the HTML template
type PageData struct {
	Title   string
	Message string
	Counter int
}

var (
	counter int
	mu      sync.Mutex
)

// homeHandler serves the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{Title: "Go + HTMX Example", Message: "Welcome to the Go + HTMX application!", Counter: counter}
	tmpl.Execute(w, data)
}

// formHandler handles form submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		message := r.FormValue("message")
		data := PageData{Title: "Go + HTMX Example", Message: message}
		tmpl := template.Must(template.ParseFiles("partials/message.html"))
		tmpl.Execute(w, data)
	}
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	mu.Unlock()
	data := PageData{Counter: counter}
	tmpl := template.Must(template.ParseFiles("partials/counter.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/submit", formHandler)
	http.HandleFunc("/increment", incrementHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
