package main

import (
	"log"
	"net"
	"bufio"
	"fmt"
	"strings"
	"time"
)

func main()  {

	li, err := net.Listen("tcp", ":8080")
	HandleErrors(err)
	defer li.Close()

	for  {
		conn, err := li.Accept()
		HandleErrors(err)

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
		fmt.Println(ln)
		if i == 0{
			//Request Line
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
			fmt.Print(conn.SetReadDeadline(time.Now().Add(5 * time.Second)))
			time.Sleep(3 * time.Second)
fmt.Print("hello")
		}
		if ln == ""{
			fmt.Print("breaking")
			break
		}
		i++
	}
}

func Respond(conn net.Conn)  {

	body := `<!DOCTYPE html><html lang="en"><head><title>Some title</title><body><strong>Hello</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 20 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func HandleErrors(err error)  {
	if err != nil{
		log.Println(err)
	}
}
