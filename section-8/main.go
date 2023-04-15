package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rabocse/lenslocked/controllers"
	"github.com/rabocse/lenslocked/templates"
	"github.com/rabocse/lenslocked/views"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS,
		"home.gohtml",
		"tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS,
		"contact.gohtml",
		"tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS,
		"faq.gohtml",
		"tailwind.gohtml"))))

	// the signup resource is a user action, therefore it might makes sense to move it
	// to a users controller package, instead of using the static.go (staticHandler).
	// This does not produce any visual difference
	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS,
		"signup.gohtml",
		"tailwind.gohtml"))

	r.Get("/signup", usersC.New)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
