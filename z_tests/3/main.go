package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main()  {

	li, err := net.Listen("tcp", ":8080")
	HandleErrors(err)

	for {
		conn, err := li.Accept()
		HandleErrors(err)

		go Handle(conn)
	}

}

func Handle(conn net.Conn)  {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "You said %s\n", ln)
	}
	defer conn.Close()
}

func HandleErrors(err error)  {
	log.Print(err)
}
