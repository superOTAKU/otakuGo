package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("server started!")
	for {
		conn, err := l.Accept()
		if err != nil {
			break
		}
		go handleRequest(&conn)
	}
}

func handleRequest(conn *net.Conn) {
	c := *conn
	data, err := ioutil.ReadAll(c)
	if err != nil {
		panic(err)
	}
	fmt.Println("from client[" + c.RemoteAddr().String() + "]: " + string(data))
	c.Write(data)
	c.Close()
}
