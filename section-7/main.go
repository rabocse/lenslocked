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

	// ! Example of rendering first the home.gohtml and then the layout-parts.gohtml.
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))))

	// ! Example of rendering first the layout-page.gohtml and then the individual contact.gohtml page.
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "contact.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
