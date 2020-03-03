package main

import (
	"fmt"
	"net"
)

type Proto struct {
	a   string
	b   net.IP
	c   int
	txt string
}

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

type RealClient struct{}

type MyClientServices interface {
	Dial(n, addr string) (net.Conn, error)
	//	Close(*net.Conn)
}

type MyClient struct {
	mClientServices MyClientServices
}

func (r RealClient) Dial(n string, addr string) (net.Conn, error) {
	fmt.Printf("Call Dial() for type %T\n", r)
	conn, err := net.Dial(n, addr)

	return conn, err
}

func (r RealClient) Close(conn *net.Conn) {
	fmt.Printf("Call Close() for type %T\n", r)

	_ = (*conn).Close()
}

func (r RealClient) PipToProto(pip Pip) *Proto {
	fmt.Printf("Call PipToProto() for type %T\n", r)
	p := Proto{}
	p.txt = "hello world"
	p.c = 0
	return &p
}

func (c MyClient) startService(pip Pip, dest string) error {
	//Connect udp
	//conn, err := net.Dial("udp", dest)
	conn, err := c.mClientServices.Dial("udp", dest)
	if err != nil {
		return err
	}
	//defer c.mClientServices.Close(&conn)

	//simple write
	conn.Write([]byte("Hello from client"))
	//todo PipToProto(pip).send(conn)

	return nil
}

func main() {
	p := "127.0.0.1:6789"
	pip := Pip{1, 2, "terrier"}

	fmt.Printf("[-] connecting to port > %s\n", p)
	//err := myclient(pip, p)
	c := RealClient{}
	srvc := MyClient{c}
	err := srvc.startService(pip, p)

	if err != nil {
		panic(err)
	}
	fmt.Println("[-] returned from connection")
}
