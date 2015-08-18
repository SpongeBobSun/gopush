// ServerInfoPush
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func ServerInfoPush(ClientAddr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ClientAddr)
	CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err)

	defer conn.Close()

	log.Println("connect success!")

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
