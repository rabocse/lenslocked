package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Alex Rabocse",
		Bio:  `<script>alert("You have been hacked!! XD") </script>`,
	}

	t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
