package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	//Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	var c net.Conn
	c, err = l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	b := make([]byte, 4)
	_, err = c.Read(b)
	if err != nil {
		log.Fatal("Failed to Read from Connection: ", err.Error())
	}
	_, err = c.Write(b)
	if err != nil {
		log.Fatal("Failed to Write to Connection: ", err.Error())
	}
	err = c.Close()
	if err != nil {
		log.Fatal("closed ", err.Error())
	}
}
