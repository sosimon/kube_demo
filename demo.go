package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	h, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Hello!\n  Hostname is: %s\n  Time is: %s\n", string(h), time.Now().Format(time.RFC1123))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running on 8080...")
	http.ListenAndServe(":8080", nil)
}
