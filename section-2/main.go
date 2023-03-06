package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my site!!!! </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>  Contact Page </h1><p> To get in touch, email me at <a href=\"mailto:example@example.test\"> exampke@example.test </a>.")
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {

	case "/":

		homeHandler(w, r)

	case "/contact":

		contactHandler(w, r)

	default:

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page Not Found")
	}

}

func main() {

	var router Router

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)

}
