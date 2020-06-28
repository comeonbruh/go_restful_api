package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Roll     string `json:"roll"`
	Password string `json:"password"`
}

type Students []Student

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	students := Students{
		Student{Id: 1, Name: "Faisal", Email: "faisal@gmail.com", Roll: "34", Password: "lkfdlakdj"},
		Student{Id: 2, Name: "Hridoy", Email: "hirdoy@gmail.com", Roll: "35", Password: "lkgslkja"},
		Student{Id: 3, Name: "Shourab", Email: "shourab@gmail.com", Roll: "36", Password: "lkgslkja"},
		Student{Id: 4, Name: "Mezbah", Email: "hirdoy@gmail.com", Roll: "62", Password: "lkgslkja"},
	}

	err := json.NewEncoder(w).Encode(students)
	if err != nil {
		fmt.Println(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit end point!")
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/students", getAllStudents).Methods("GET")
	log.Fatal(http.ListenAndServe(":9090", r))
}

func main() {
	handleRequests()
}
