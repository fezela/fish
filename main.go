package main

import (
	"fmt"
	"log"
	"net/http"
	"fish/kanban"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const connectionString = "postgres://fish:development@postgres:5432/fishdb?sslmode=disable"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go behind Apache SS!")
}

func radio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Radio page from Go gateway")
}

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Api routing in development")
}

func main(){
	db, err := sql.Open("pgx", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to PostgreSQL")
	
	var databaseName string
	dbErr := db.QueryRow("SELECT current_database()").Scan(&databaseName) //& operator indicates address in memory

	if dbErr != nil {
		log.Fatal(dbErr)
	}
	log.Println("Connected to database: ", databaseName)

	http.HandleFunc("/", home)
	http.HandleFunc("/radio", radio)
	http.HandleFunc("/api", api)
	http.HandleFunc("/kanban", kanban.Handler)

	log.Println("Go gateway listening on http://localhost:8080")
	httpErr := http.ListenAndServe("0.0.0.0:8080", nil)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
