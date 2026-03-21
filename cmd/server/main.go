package main

// just fyw InitDb() in db.go does not need to be imported like a module in js land
import (
	"fmt"
	"os"

	"github.com/CPoooo/goHttpServer/internal/db"
)

func main() {
	// conn here is a dbpool - name it that?
	dbpool, err := db.InitDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to DB: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	// do not defer the closing of the db connection from inside the init to the db! (duh)
	// defer dbpool.Close()
} // defer running btw -- closing connection conn.Close
