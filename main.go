package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome website</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1> <p>To get in touch, email me at <a href=\"mailto:contact@gamil.com\">contact@gmail.com</a>.")
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

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "whoops, 404 it is, that's no gouda, try another path ", http.StatusNotFound)
	}
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "whoops, 404 it is, that's no gouda, try another path ", http.StatusNotFound)
// 	}
// }

func main() {
	fmt.Println("Starting server on : 3000...")
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
