package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const josnFilePath = "./users.json"

func main() {
	content, err := os.ReadFile(josnFilePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully opened users.json")

	var users users
	err = json.Unmarshal(content, &users)
	if err != nil {
		panic(err)
	}

	for _, user := range users.Users {
		user.PrintInfo()
	}

	fmt.Println(users)
	os.Exit(0)
}

type users struct {
	Users []user `json:"users`
}

type user struct {
	Name   *string `json:"name"`
	Type   *string `json:"type"`
	Age    uint8   `json:"age"`
	Social struct {
		Facebook *string `json:"facebook"`
		Twitter  *string `json:"twiiter"`
	} `json:"social"`
}

func (u user) PrintInfo() {
	fmt.Println("User Type: ", u.Type)
	fmt.Println("User Age: ", u.Age)
	fmt.Println("User Name: ", u.Name)
	fmt.Println("Facebook Url: ", u.Social.Facebook)
}
