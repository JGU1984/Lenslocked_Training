package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//----- OS agnostic way, this to set / or \ correctly -----
	// tplPath := filepath.Join("templates", "home.gohtml")
	// tpl, err := template.ParseFiles(tplPath)
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Clean("templates/home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Clean("templates/contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ</h1> 
	<ul> 
		<li>Q: What are your support hours?</li>
		<li>A: Yes! We offer a free trail for 30 days on any paid plans.</li><br>
		<li>Q: What are your support hours?</li>
		<li>A: We have support staff answering emails 24/7, though reponse times may be a bit slower on weekends.</li><br>
		<li>Q: How do I contact support?</li>
		<li>A: Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a></li>
	</ul>`)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "whoops, 404 it is, that's no gouda, try another path ", http.StatusNotFound)
	})

	fmt.Println("Starting server on : 3000...")
	http.ListenAndServe(":3000", r)
}
