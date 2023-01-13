package main

import (
	"fmt"
	"log"
	"net/http"
)

func welcomehandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[9:] == "hello" {
		fmt.Fprintf(w, "Hey, %s!", r.URL.Path[9:])
	} else {
		fmt.Fprintf(w, "Hello %s!, Welcome here!!", r.URL.Path[9:])
	}
}
func nameshandler(w http.ResponseWriter, r *http.Request) {

	guys, ok := r.URL.Query()["name"]

	if !ok || len(guys[0]) < 1 {
		fmt.Fprintf(w, "Hayk Baluyan\nTeja Swaroop\nPuneeth Sharma")
		return
	} else if guys[0] == "Puneeth" || guys[0] == "Teja" || guys[0] == "Hayk" {
		fmt.Fprintf(w, "Hi %s!!", guys[0])
	} else {
		fmt.Fprintf(w, "Unknown Name")
	}
}

// func specifichandler(w1 http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path[27:] == "Hayk" || r.URL.Path[27:] == "Puneeth" || r.URL.Path[27:] == "Teja" {
// 		fmt.Fprintf(w1, "Hello %s!!", r.URL.Path[27:])
// 	} else {
// 		fmt.Fprintf(w1, "Hello Unknown!!")
// 	}

// }
func main() {
	http.HandleFunc("/averlon/", welcomehandler)
	http.HandleFunc("/averlon/names/", nameshandler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
