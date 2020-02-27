package main

import (
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
