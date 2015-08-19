// server.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	//	"strconv"
	//	"time"
)

func main() {
	//number of core
	var MULTICORE int = runtime.NumCPU() //number of core
	log.Println("CORE_NUM:", MULTICORE)
	runtime.GOMAXPROCS(MULTICORE)
	//Listen the port of Server
	netListener, err := net.Listen("tcp", ":22333")
	CheckError(err)
	defer netListener.Close()

	log.Println("Waiting for the Clients")
	for {
		conn, err := netListener.Accept()
		if err != nil {
			continue
		}
		log.Println("ClientAddr:", conn.RemoteAddr().String())
		log.Println("Msg:connect to client success by tcp ")
		go handleConnection(conn) //此处使用go关键字新建线程处理连接，实现并发
		//from server to clients
		//		ServerInfoPush(conn.RemoteAddr().String()) //
		//		TcpSender(conn)
	}
	//	TcpSender(conn)

	fmt.Println("Hello World!")
}
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		//		log.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		//		 n, err := conn.Read(buffer)
		l, err := conn.Read(buffer)
		if err != nil {
			log.Println(conn.RemoteAddr().String(), "connection error: ", err)
			return
		}
		log.Println(conn.RemoteAddr().String(), "receive data length:", l)
		log.Println(conn.RemoteAddr().String(), "receive data: ", buffer[:l])
		log.Println(conn.RemoteAddr().String(), "receive data string: ", string(buffer[:l]))
		go loopingCall(conn)
	}

}
