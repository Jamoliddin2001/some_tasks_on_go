package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Professor struct {
	Name       string     `json:"name"`
	ScienceID  int        `json:"science_id"`
	IsWorking  bool       `json:"is_Working"`
	University University `json:"university"`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	prof1 := Professor{
		Name:      "Bob",
		ScienceID: 1,
		IsWorking: true,
		University: University{
			Name: "MSU",
			City: "Moscow",
		},
	}

	byteArr, err := json.MarshalIndent(prof1, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteArr))

	err = os.WriteFile("./course/lec1/murshal/output.json", byteArr, 777)
	if err != nil {
		log.Fatal(err)
	}
}
