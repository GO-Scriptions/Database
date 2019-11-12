package main

import (
	"fmt"
	"net/http"
)

var mess = "Hello from the database."

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/c", cu)

	http.ListenAndServe(":80", nil)
}

func cu(w http.ResponseWriter, q *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Println(nickIPs[ip], "has joined the chat room!")
	buffer.Execute(w, q)
}
