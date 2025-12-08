package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	main1()
}



func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main1() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
