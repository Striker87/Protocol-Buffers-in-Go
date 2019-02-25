package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	user := &Person{
		Name: "Gendos",
		Age: 32,
		SocialFollowers: &SocialFollowers{
			Twitter: 1400,
			Youtube: 2000,
		},
	}

	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	newUser := &Person{}
	err = proto.Unmarshal(data, newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newUser.GetAge())
	fmt.Println(newUser.GetName())
	fmt.Println(newUser.SocialFollowers.GetYoutube())
	fmt.Println(newUser.SocialFollowers.GetYoutube())
}
