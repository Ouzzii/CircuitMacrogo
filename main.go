package main

import (
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Username string
	Fname    string
}

func main() {

	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func h1(w http.ResponseWriter, r *http.Request) {
	temPlate := template.Must(template.ParseFiles("index.html"))
	Users := []User{
		{Username: "Doe", Fname: "John"},
		{Username: "0zaninyo", Fname: "Ozan"},
	}

	temPlate.Execute(w, Users)
}
