package main

type Users struct {
	Users []User
}

type User struct {
	name   string `json:"name"`
	age    int    `json:"age"`
	friend Friend `json:"friend"`
}

type Friend struct {
	name string
}
