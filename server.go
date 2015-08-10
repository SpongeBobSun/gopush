// server
package main

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
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
		fmt.Println("host:", req.Host)
		fmt.Println("header:", req.Header)
		fmt.Println("RequestURI", req.RequestURI)
		fmt.Println("URL", req.URL)
		fmt.Println("Body", req.Body)
		fmt.Println("req.RemoteAddr", req.RemoteAddr)
		UUIDKey := uuid.New()
		r.JSON(200, map[string]interface{}{"service_id": service_id, "key": UUIDKey})
		//		conn,err := http.

	})

	fmt.Println("Hello World!")
	m.RunOnAddr(":1111")
	http.Handle("/", m)
}
