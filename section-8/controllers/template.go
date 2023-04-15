package controllers

import "net/http"

/*

if we check the users.go controller:

type Users struct {
	Templates struct {
		New views.Template
	}
}

We can see that the views.Template is used. Which mean that we are locked in (forced) to use that package views.

If we replace views.Template with an interface, we could pass anything as long as it satisfies the interface. In other words, as long as
it passes the methods the interface needs, then it will work.

What are that (those) methods?

At this point, if we check the users.go, we can see that the only method being used is "Execute" method:

func (u Users) New(w http.ResponseWriter, r *http.Request) {

	u.Templates.New.Execute(w, nil)

}

And, if we check static.go, we will notice that staticHandler, and FAQ are also only requiring the "Execute" method.

Therefore, the interface only needs the "Execute" method to be satisfied.

*/

type Template interface {

	// data is the interface used to execute the template.
	Execute(w http.ResponseWriter, data interface{})
}
