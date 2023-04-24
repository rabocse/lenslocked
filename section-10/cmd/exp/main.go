package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // stdlib is the package that provides the standard library driver. To be able to use database sql package
)

func main() {

	// pgx = Driver name defined by github.com/jackc/pgx/v4/ It we were using a different library, the driver would  be different.
	// host, user, password, dbname were the values used in the YAML file for the docker compose.
	db, err := sql.Open("pgx", "host=localhost port=5432 user=baloo password=junglebook dbname=lenslocked sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

}
