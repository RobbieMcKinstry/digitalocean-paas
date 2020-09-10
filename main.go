package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Running on port :80")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
