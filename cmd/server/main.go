package main

import "fmt"

func main() {
	fmt.Println("connecting to db soon...")
}

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
