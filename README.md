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
By Jean.Jiang

* Tcp connection established successfully from server to client.

* Create a buffer([]byte,50) to receive tcp data.

* The author of the Client is too LAAAAAAAZY (?)

###2015-08-18:
By Jean.Jiang

* Modify to push data from server to client

* ~~Now I fall in some question of the information push,After I will kill them FAST.~~ WTF does that sentence means ??

* A Wonderful Day

###2015-08-19:
By Jean.Jiang

* First,I kill the question of ~~informationpush~~ (Do you mean message pushing ?)

* Add a `loopingCall` function to send ~~informationpush(json)~~(data?) to clients

		func loopingCall(conn net.Conn) {
			pingTicker := time.NewTicker(10 * time.Second) //set timer
			testAfter := time.After(2 * time.Second)       //set timer delay
			for {
				select {
				case <-pingTicker.C:
					//send heartbeat
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
		   send data
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


###2015-08-19
By Bob.Sun
Typo correction.

Grammar check.

Actually I DO NOT think we should log our work in this section.

And, I'm being lazy because client side code is stable.

## Credits
[Jean.Jiang](https://github.com/JiangXuanYi)

[Bob.Sun](https://github.com/SpongeBobSun)




