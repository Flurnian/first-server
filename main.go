package main

import (
	"fmt"
	"log"
	"net/http"
)

// todo handle only GET REQUEST. Use some library
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, Vadym"))
	})

	port := 8090
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil)) //:8090
}
