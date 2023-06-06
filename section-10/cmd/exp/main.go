package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // stdlib is the package that provides the standard library driver. To be able to use database sql package
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {

	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	// pgx = Driver name defined by github.com/jackc/pgx/v4/ It we were using a different library, the driver would  be different.
	// host, user, password, dbname were the values used in the YAML file for the docker compose.
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	// Create a table

	_, err = db.Exec(
		`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount INT,
		description TEXT
	);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created")

	// name := "Dar Sal"
	// email := "dar@fakemail.com"

	// _, err = db.Exec(
	// 	`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2);`, name, email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created")

	// //Adding another user for testing purposes (same than above attempt)

	// name2 := "Bastian Escobar"
	// email2 := "bastian@fakemail.com"

	// _, err = db.Exec(
	// 	`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2);`, name2, email2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created")

	// Insert more data

	newName := "Lionel"
	newEmail := "lmessi@thegoat.example"

	row := db.QueryRow(`
	INSERT INTO users (name, email)
	VALUES ($1, $2) RETURNING ID;`, newName, newEmail)
	var id int
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("User created id=", id)

}
