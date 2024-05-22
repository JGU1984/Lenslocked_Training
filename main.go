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

	tpl := views.Must(views.Parse(filepath.Clean("templates/home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Clean("templates/contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Clean("templates/faq.gohtml")))))

	tpl = views.Must(views.Parse(filepath.Clean("templates/info.gohtml")))
	r.With(controllers.InfoDataMiddleware).Get("/info", controllers.StaticHandlerInfo(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "whoops, 404 it is, that's no gouda, try another path ", http.StatusNotFound)
	})

	fmt.Println("Starting server on : 3000...")
	http.ListenAndServe(":3000", r)
}
