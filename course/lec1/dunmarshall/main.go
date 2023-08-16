package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	jsonFile, err := os.Open("/home/jamoliddin/GolandProjects/lesson1/course/lec1/unmarshall/users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("Success")

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	fmt.Println(result)
}
