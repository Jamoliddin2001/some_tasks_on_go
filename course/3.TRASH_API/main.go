package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	port string = "8080"
)

func main() {
	log.Println("Trying to start REST API pizza!")
	// Router
	router := mux.NewRouter()
	router.HandleFunc("/pizza", getAllPizza).Methods("GET")
	router.HandleFunc("/pizza/{id}", getPizzaById).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getAllPizza(writer http.ResponseWriter, request *http.Request) {

}

func getPizzaById(writer http.ResponseWriter, request *http.Request) {

}
