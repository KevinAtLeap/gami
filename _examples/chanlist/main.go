package main

import (
	"log"
	"time"

	"github.com/CyCoreSystems/gami"
)

func main() {
	ami, err := gami.Dial("127.0.0.1:5038")
	if err != nil {
		log.Fatalln("failed to connect to AMI:", err)
	}
	ami.Run()
	defer ami.Close()

	go func() {
		for {
			select {
			case err := <-ami.NetError:
				log.Println("network error:", err)
				<-time.After(time.Second)
				if err := ami.Reconnect(); err == nil {
					ami.Action("Events", gami.Params{
						"EventMask": "on",
					})
				}
			case err := <-ami.Error:
				log.Println("generic error:", err)
			case ev := <-ami.Events:
				log.Println("event: ", ev)
			}
		}
	}()

	if err := ami.Login("admin", "testmenow"); err != nil {
		log.Fatalln("login failed:", err)
	}

	if _, err := ami.Action("Ping", nil); err != nil {
		log.Fatalln("ping failed:", err)
	}

	resp, err := ami.Action("CoreShowChannels", nil)
	if err != nil {
		log.Fatalln("failed to get list of channels:", err)
	}
	for k, v := range resp.Params {
		log.Printf("%s: %s", k, v)
	}
}
