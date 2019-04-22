package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	echoHandler := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Echo response")
	}

	fmt.Println("Hello Go!")
	fmt.Println("Arguments:", os.Args)
	// fmt.Println("Environments:", os.Environ())
	http.HandleFunc("/echo", echoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
