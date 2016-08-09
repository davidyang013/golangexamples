package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println(conn.LocalAddr())

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

	tcpAdd, _ := net.ResolveTCPAddr("tcp4", "baidu.com:80")
	fmt.Println(tcpAdd)
}
