package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var name, _ = os.Hostname()
	fmt.Fprintf(w, "Hello world, I'm running on %s. Build Number %s.\n", name, os.Getenv("BUILDNUMBER"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
