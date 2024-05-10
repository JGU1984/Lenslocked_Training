package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/jgu1984/lenslocked/controllers"
	"github.com/jgu1984/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	tpl, err := views.Parse(filepath.Clean("templates/home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Clean("templates/contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Clean("templates/faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "whoops, 404 it is, that's no gouda, try another path ", http.StatusNotFound)
	})

	fmt.Println("Starting server on : 3000...")
	http.ListenAndServe(":3000", r)
}
