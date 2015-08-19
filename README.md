# GoPush
## Introducton

GoPush is a message pushing library for Android devices and was inspired by [NoPush](https://github.com/SpongeBobSun/NoPush)

The Android client basicly is the same with NoPush and the server side is written by GoLang so we get a better concurrrent performance.

Also, GoPush is intent to create a free & open-source message pushing service.

GoPush can be deployed on your own server and work along with your own web applications.

## Project Status
Under developing and testing.

## Usage
## Server Interface
## Client Interface
## Design (How does it works?)
###2015-08-17:
By [Jean.Jiang](https://github.com/JiangXuanYi)

1,Tcp connect is success from server to client.

2,Create a buffer([]byte,50) to receive the tcp information

3,The author of the Client is too LAAAAAAAZY
###2015-08-18:
By [Jean.Jiang](https://github.com/JiangXuanYi)

1,Add a serverclient to push the information(json),from server to clients

2,Now I fall in some question of the informationpush,After I will kill them FAST.

3,A Wonderful Day
###2015-08-19:
By [Jean.Jiang](https://github.com/JiangXuanYi)

1,First,I kill the question of informationpush

2,Add a loopingCall to send the informationpush(json) to the clients

```GoLang
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

```



## Credits
[Jean.Jiang](https://github.com/JiangXuanYi)

[Bob.Sun](https://github.com/SpongeBobSun)




