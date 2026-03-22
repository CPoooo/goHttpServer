package main

// just fyw InitDb() in db.go does not need to be imported like a module in js land
// know how import works and how go looks for/uses packages
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/CPoooo/goHttpServer/internal/db"
)

/*
Use bufio to read not a "raw" buffer (will buffer the information that we get from conn.Read)
http is delimeted by  "\r\n"

Every time you call conn.Read() directly that's a syscall — you're asking the OS to go get data from the network.
Syscalls are expensive relative to reading from memory.

bufio reads ahead in 4096 byte chunks internally (from that syscall to get data from the network - conn.Read())
ReadString('\n') serves from memory, not from kernel each time
this is why bufio exists -- batch syscalls, not one per line
*/
// what about query params? and dynamic urls?
// pick one from verb
func GET(url string) []byte {
	fmt.Printf("GET request made to %v", url)
	// just send back a 200 ok res for now
	s := "HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n"
	return []byte(s)
}
func POST(url string, body RequestBody) []byte {
	fmt.Printf("POST request made to %v with body: %v", url, body)
	s := "HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n"
	return []byte(s)
}
func PUT(url string, body RequestBody) []byte {
	fmt.Printf("PUT request made to %v with body: %v", url, body)
	s := "HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n"
	return []byte(s)
}

// delete will always have a /:id or something (unless we give admin ability to DELETE all /users)
func DELETE(url string) []byte {
	fmt.Printf("GET request made to %v", url)
	s := "HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n"
	return []byte(s)
}

type RequestBody []byte // i guess idk (this will represent the array of bytes coming from http body section)
// do we need to attach a content-length to http methods like POST() that require a body
// whatever data a http response will need
type Response struct {
	body string // i mean this should really be a JSON type or something like that or just []byte or *byte idk
}

// whatever data a http request will have (idk - i guess i am building a lib here)
type Request struct {
	verb   string // enum? iota in Go
	url    string
	header map[string]string
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// so conn here was/is/has an io.Reader()? for bufio.NewReader() to read from
	// reading from our reader that is reading from conn write?
	reader := bufio.NewReader(conn) // reader: *bufio.Reader

	// fmt.Printf("New connection received: %v\n", conn)
	// buffer := make([]byte, 4096)

	// 1. start line
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading line: %v\n", err)
	}

	splitLine := strings.Split(line, " ")
	verb := splitLine[0]
	url := splitLine[1]
	protocol := splitLine[2]
	body := []byte{0x68, 0x65, 0x79}
	fmt.Println(line)
	fmt.Printf("Verb: %v\nUrl: %v\nProtocol: %v\n", verb, url, protocol)
	headers := map[string]string{}

	// 2. headers
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading line: %v\n", err)
			break
		}

		// 3. empty line
		if line == "\r\n" {
			// everything after this point should be the body
			break
		}
		headerSplit := strings.SplitN(line, ":", 2)
		key := strings.Trim(headerSplit[0], " ")
		value := strings.Trim(headerSplit[1], " ")
		headers[key] = value
	}
	// 4. body
	response := []byte{}
	switch verb {
	case "GET":
		response = GET(url)
	case "POST":
		response = POST(url, body)
	case "PUT":
		response = PUT(url, body)
	case "DELETE":
		response = DELETE(url)
	}
	n, err := conn.Write(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing response %v\n", err)
	}
	fmt.Printf("%v bytes written", n)
}

// notes for yourself:
// THIS READS FROM THE STREAM SO THERE IS NO STREAM LATER TO LOOP OVER
// n, err := reader.Read(buffer)
// if err != nil {
// 	fmt.Fprintf(os.Stderr, "Failed to read from TCP Connection: %v\n", err)
// 	os.Exit(1)
// }
// fmt.Printf("Received Request:\n%s\n", buffer[:n])

type User struct {
	name string
	age  int
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
	fmt.Printf("Listening for TCP Connections on http://localhost:8080\n")
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
