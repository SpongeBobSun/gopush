// client_test
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	//	serverAddr := "127.0.0.1:22333"
	//	serverAddr := "10.10.14.231:22333"
	serverAddr := "10.10.14.231:60534"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	checkError(err)
	defer conn.Close()

	fmt.Println(" connect success!")

	go TcpSender(conn)
	for {
		time.Sleep(1 * 1e9)
	}
	fmt.Println("Hello World!")
}

func TcpSender(conn net.Conn) {
	msg := "{\"notification\":\"notification\",\"message\":\"message\", \"icon\": \"http://www.baidu.com\"}"
	conn.Write([]byte(msg))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
