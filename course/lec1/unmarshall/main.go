package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []Friend `json:"friends"`
}

type Friend struct {
	Name string `json:"name"`
}

func printUser(u *User) {
	fmt.Println("Name", u.Name)
	fmt.Println("Age", u.Age)
	fmt.Println("Friends", u.Friends)
}

func main() {
	jsonFile, err := os.Open("/home/jamoliddin/GolandProjects/lesson1/course/lec1/unmarshall/users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("Success")
	var users Users

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &users)
	for _, value := range users.Users {
		printUser(&value)
	}
}
