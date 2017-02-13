package main

import (
	"net"
	"log"
	"io"
	"fmt"
)

func main()  {

	//Allows for connections to come through the TCP layer on port :8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Panic(err)
	}
	defer li.Close()

	for {
		//Accept() waits for and returns the next connection to the listener.
		conn, err  := li.Accept()
		if err!= nil{
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
