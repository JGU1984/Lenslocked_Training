package controllers

import (
	"context"
	"net/http"

	"github.com/jgu1984/lenslocked/views"
)

type ctxKey string

const infoKey ctxKey = "info"

type Info struct {
	Name  string
	Age   int
	Learn string
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func StaticHandlerInfo(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, ok := r.Context().Value(infoKey).(Info)
		if !ok {
			http.Error(w, "No info data found", http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, info)
	}
}

func InfoDataMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/info" {
			info := Info{
				Name:  "Johan",
				Age:   39,
				Learn: "Go",
			}
			ctx := context.WithValue(r.Context(), infoKey, info)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
