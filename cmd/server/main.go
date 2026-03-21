package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, db)
	// still need to learn a Context fully here but I know that it is something to do
	// with Go attaching "context" to and from client/server communications in go (more research to be done here)
	// now this conn here will be our '*pgx.conn' type - connection to db
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// pushing conn.Close to defer stack
	defer conn.Close(context.Background())

	var greeting string
	// Looks to just run a hello world query - select 'hello, world!' - what does this do in sql btw
	// Scan returns this greeting into &greeting
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
} // defer running btw -- closing connection conn.Close

/* TODO:
goal min: crud api
GET    /users
GET    /users/{id}/todos
POST   /todos
PATCH  /todos/{id}
DELETE /todos/{id}

connect to db

open up socket and listen for connections
parse req and answer with json

error handling bad requests

auth?
*/
