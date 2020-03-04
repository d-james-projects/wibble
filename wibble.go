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

/*
func NewPip(t string) *Pip {
	p := Pip{txt: t}
	p.x = 0
	p.y = 0
	return &p
}
*/
type RealClient struct{}

type MyClientServices interface {
	Dial(n, addr string) (net.Conn, error)
	//Close(*net.Conn)
	PipToProto(pip Pip) *Proto
}

type Sender interface {
	Send(conn *net.Conn) error
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
	p := new(Proto)
	p.txt = "hello world"
	p.c = 0
	return p
}

func (p *Proto) Send(conn *net.Conn) error {
	fmt.Printf("Call Send() for type %T\n", p)
	fmt.Println("Make a real call to the send func")
	_, err := net.ResolveUDPAddr("realnet", "realaddress")
	return err
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
	//not called
	//conn.Write([]byte("Hello from client"))
	c.mClientServices.PipToProto(pip).Send(&conn)

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
