package main

import (
	"fmt"
	"net"
)

func myclient(dest string) error {
	//Connect udp
	conn, err := net.Dial("udp", dest)
	if err != nil {
		return err
	}
	defer conn.Close()

	//simple write
	conn.Write([]byte("Hello from client"))

	return nil
}

func main() {
	p := "127.0.0.1:6789"

	fmt.Printf("[-] connecting to port > %s\n", p)
	err := myclient(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("[-] returned from connection")
}
