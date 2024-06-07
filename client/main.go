package main

import (
	"log"
	"net/rpc"
	"taqsir/utlis"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	user, err := utlis.ReadFromFile("user.json")
	if err != nil {
		log.Fatal("error user reading: ", err)
	}

	for _, u := range user {
		var reply string
		err = client.Call("UserHadler.CreateUser", u, &reply)
		if err != nil {
			log.Fatal("error creating user: ", err)
		} else {
			log.Println("user created: ", reply)
		}
	}
}
