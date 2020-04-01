package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

/**
go 的两个基本接口：io.Reader, io.Writer，相当于java的InputStream， OutputStream
*/

func TestReader(t *testing.T) {
	reader := strings.NewReader("hello world")
	buf := make([]byte, 5)
	len, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:len]))
}

func TestWriter(t *testing.T) {
	//类似Java的ByteArrayOutputStream，但同时实现了Reader接口
	buf := bytes.Buffer{}
	buf.Write([]byte("hello"))
	fmt.Println(buf.String())
}

/**
go对文件的操作更底层，并且File本身实现了Reader和Writer接口，不需要额外操作
*/
func TestReadWriteFile(t *testing.T) {
	const tmpFile = "/tmp/notexist"
	f, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}
	f.Close()
	f, err = os.OpenFile(tmpFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 5)
	if len, err := f.Read(buf); err != nil {
		panic(err)
	} else {
		fmt.Println(string(buf[:len]))
	}
	f.Close()
	os.Remove(f.Name())
}

/**
buffer io，相当于java的 BufferInputStream BufferOutputStream
*/
func TestBuffer(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 100))
	bufWriter := bufio.NewWriter(buf)
	bufWriter.Write([]byte("hello"))
	bufWriter.Write([]byte(" world"))
	bufWriter.Flush()
	fmt.Println(buf.String())
}
