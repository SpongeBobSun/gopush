// server.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
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
		log.Println(conn.RemoteAddr().String(), "connect to client success by tcp ")
		go handleConnection(conn)
		//from server to clients
		//		ServerInfoPush(conn.RemoteAddr().String())//

	}

	fmt.Println("Hello World!")
}
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
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
	}

}
