package main

// just fyw InitDb() in db.go does not need to be imported like a module in js land
// know how import works and how go looks for/uses packages
import (
	"fmt"
	"net"
	"os"

	"github.com/CPoooo/goHttpServer/internal/db"
)

func handleConnection(conn net.Conn) {
	fmt.Printf("New connection received: %v\n", conn)
}

func main() {
	// conn here is a dbpool - name it that?
	dbpool, err := db.InitDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to DB: %v\n", err)
		os.Exit(1)
	}
	// do not defer the closing of the db connection from inside the init to the db! (duh)
	defer dbpool.Close()

	// open tcp
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create TCP Listener: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()
	for {
		// Read bytes from tcp stream
		connection, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to Accept() on TCP Listener: %v\n", err)
			os.Exit(1)
		}
		go handleConnection(connection)
	}

	// get request

	// print request

	// parse and marshal/unmarshal to and from json - how to know if the bytes from req
	// that i received are http vs another protocol

	// marshal into json and send back over tcp to client
	// each connection needs a goroutine handler to be run

}
