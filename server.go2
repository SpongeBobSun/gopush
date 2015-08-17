// server
package main

import (
	//	"time"
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	//	"net"
	//	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func main() {

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/push", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		service_id := req.FormValue("service_id")
		fmt.Println("req.RemoteAddr", req.RemoteAddr)
		//1：收到请求连接之后， 记录客户端的service_id,如果此service_id在数据库中有记录，则取出 key值，并返回给客户端
		//2：同时，记录初次连接时间，
		//3：建立tcp 长链接心跳，初步，tcp心跳周期为1min

		//4：心跳存在时，可以在某些时候，给客户端，推送信息。
		//最后，去除Martini开元框架，用REST自定义框架，代替
		//		conn, err := net.Dial("tcp", req.RemoteAddr)
		//		if err != nil {
		//			panic(err)
		//		}

		//		fmt.Println("conn", conn)
		UUIDKey := uuid.New()
		r.JSON(200, map[string]interface{}{"service_id": service_id, "key": UUIDKey})
		//		conn,err := http.

	})

	fmt.Println("Hello World!")
	m.RunOnAddr(":1111")
	http.Handle("/", m)
}
