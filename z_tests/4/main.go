package main

import (
	"log"
	"net"
	"bufio"
	"fmt"
	"strings"
)

func main()  {

	li, err := net.Listen("tcp", ":8080")
	HandleError(err)
	defer  li.Close()

	for{
		conn, err := li.Accept()
		HandleError(err)

		go Handle(conn)
	}
}

func Handle(conn net.Conn)  {

	defer conn.Close()

	Request(conn)

	Respond(conn)

}

func Request(conn net.Conn)  {

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Print(ln)
		if i == 0{
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			fmt.Print(m)
			fmt.Print(u)
		}
		if ln ==""{
			break
		}
		i++
	}
}

func Respond(conn net.Conn)  {

	body := `<DOCTYPE html><html><title>NewTitle</title><body><h1>This is the body</html></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func HandleError(err error)  {
	if err != nil {
		log.Println(err)
	}
}


