package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	//DialTcp返回tcp连接，直接dial返回的连接只能Close，不能CloseWrite
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{net.ParseIP("localhost"), 8080, ""})
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("hello"))
	//tcp可以半关闭
	conn.CloseWrite()
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("from server[" + conn.RemoteAddr().String() + "]: " + string(data))
	conn.Close()
}
