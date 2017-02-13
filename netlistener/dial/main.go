package main

import (
	"net"
	"log"
	"io/ioutil"
	"fmt"
)

func main()  {

	conn,err := net.Dial("tcp", "localhost:8080")
	if err != nil{
		log.Println(err)
	}
	defer  conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil{
		log.Println(err)
	}

	fmt.Println(conn.LocalAddr().String())

	fmt.Println(string(bs))
}
