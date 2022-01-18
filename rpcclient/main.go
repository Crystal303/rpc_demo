package main

import "log"

func main() {
	client, err := NewHelloServiceClient("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("client init err: ", err)
	}
	var reply string
	if err := client.Hello("lao ang", &reply); err != nil {
		log.Fatal("call hello err: ", err)
	}
	log.Print(reply)
}
