package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", start)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start")

}
