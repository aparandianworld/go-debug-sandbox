package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			name = "Aaron"
		}

		if name == "crash" {
			panic("crashing")
		}

		greeting := fmt.Sprintf("Hello, %s!", name)
		fmt.Fprint(w, greeting)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server starting on http://localhost:" + port)
	log.Println("Try: http://localhost:" + port + "/?name=Aaron")
	log.Println("Or crash it: http://localhost:" + port + "/?name=crash")

	http.ListenAndServe(":"+port, nil)
}
