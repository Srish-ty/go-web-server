package main

import (
	"fmt"
	"net/http"

	"example.com/hello/result"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, result.ResString)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port :3001")
	fmt.Println("http://localhost:3001/")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
