package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>This is my <i>Awesome WebServer</i> V2.0!</b>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening at : 8080")
	http.ListenAndServe(":8080", nil)
}
