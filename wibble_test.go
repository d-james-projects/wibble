package main

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockClient struct {
	mock.Mock
}

/*
func (m *mockClient) Close(conn *net.Conn) {
	fmt.Printf("Call Close() for type %T\n", *m)
	_ = m.Called(conn)
}
*/
func (m *mockClient) Dial(n string, addr string) (net.Conn, error) {
	fmt.Printf("Call Dial() for type %T\n", *m)
	args := m.Called(n, addr)
	return args.Get(0).(net.Conn), args.Error(1)
}

func (m *mockClient) PipToProto(pip Pip) *Proto {
	fmt.Printf("Call PipToProto() for type %T\n", *m)
	args := m.Called(pip)
	return args.Get(0).(*Proto)
}

func TestStartService(t *testing.T) {
	p := "127.0.0.1:6789"
	pip := Pip{1, 2, "terrier"}
	c := new(RealClient)
	srvc := MyClient{c}

	fmt.Println("========Call 1")

	assert.Nil(t, srvc.startService(pip, p))

	fmt.Println("========Call 2")

	assert.NotNil(t, srvc.startService(pip, "wibble"))
}

func TestStartServiceMocked(t *testing.T) {
	pip := Pip{1, 2, "terrier"}
	//test := new(net.Conn)
	obj, _ := net.Pipe()
	defer obj.Close()
	ptp := new(Proto)

	c := new(mockClient)

	c.On("Dial", "udp", "127.0.0.1:6789").Return(obj, nil).Once()
	//c.On("Dial", "udp", "wobble").Return(obj, errors.New("bad net")).Once()
	c.On("PipToProto", pip).Return(ptp).Once()

	srvc := MyClient{c}

	//module under test
	fmt.Println("========Call 3")
	srvc.startService(pip, "127.0.0.1:6789")
	//srvc.startService(pip, "wobble")

	c.AssertExpectations(t)

}

/*
func TestMyclient(t *testing.T) {

	assert.Nil(t, myclient("127.0.0.1:6789"))

	assert.NotNil(t, myclient("wibble"))

}

func server(done chan int, port string) {

	log.Println("start server ...")

	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	log.Println("listening ...")

	//simple read
	buffer := make([]byte, 1024)
	n, _, err := pc.ReadFrom(buffer)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("> %d bytes read", n)
	done <- n
}

func TestMyclientWithConn(t *testing.T) {

	log.SetOutput(os.Stdout)

	port := "127.0.0.1:7890"
	done := make(chan int, 1)
	go server(done, port)

	time.Sleep(time.Second)
	log.Println("main thread resuming ...")

	assert.Nil(t, myclient(port))

	// wait for go routine and assert bytes read
	r := <-done
	assert.Equal(t, 17, r, "should be equal")
}
*/
