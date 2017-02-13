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

	defer li.Close()

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
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Print("Code will get here")
}

func HandleErrors(err error)  {
	if err != nil{
		log.Println(err)
	}
}