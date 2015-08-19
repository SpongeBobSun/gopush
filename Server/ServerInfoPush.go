// ServerInfoPush
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func loopingCall(conn net.Conn) {
	pingTicker := time.NewTicker(10 * time.Second) //定时
	testAfter := time.After(2 * time.Second)       //延时
	for {
		select {
		case <-pingTicker.C:
			//发送心跳心跳数据
			msg := "{\"notification\":\"notification\",\"message\":\"message\", \"icon\": \"http://www.baidu.com\"}"
			_, err := sendData(conn, []byte(msg))
			if err != nil {
				pingTicker.Stop()
				return
			}
		case <-testAfter:
			log.Println("testAfer:")

		}
	}
}

/*
   发送数据
*/
func sendData(conn net.Conn, data []byte) (n int, err error) {
	addr := conn.RemoteAddr().String()
	n, err = conn.Write(data)
	if err == nil {
		//		log.println("=>", addr, string(data))
		log.Println("=>", addr, string(data))
	}
	return
}

func ServerInfoPush(ClientAddr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ClientAddr)
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
