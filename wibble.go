package main

import (
	"fmt"
	"net"
)

type Pip struct {
	x   int
	y   int
	txt string
}

func NewPip(t string) *Pip {
	p := Pip{txt: t}
	p.x = 0
	p.y = 0
	return &p
}

type TheRealClient struct{}

type MyClientServices interface {
	Dial(n string, addr string) (net.Conn, error)
	Close()
}

type MyClient struct {
	m MyClientServices
}

func myclient(pip Pip, dest string) error {
	//Connect udp
	conn, err := net.Dial("udp", dest)
	if err != nil {
		return err
	}
	defer conn.Close()

	//simple write
	//conn.Write([]byte("Hello from client"))
	//todo PipToProto(pip).send(conn)

	return nil
}

func main() {
	p := "127.0.0.1:6789"
	pip := Pip{1, 2, "terrier"}

	fmt.Printf("[-] connecting to port > %s\n", p)
	err := myclient(pip, p)
	if err != nil {
		panic(err)
	}
	fmt.Println("[-] returned from connection")
}
