package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/efnky/go-internal-pkg/internal/greeter"
	"github.com/efnky/go-internal-pkg/pkg/version"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s (v%s) up on :%s\n", greeter.Greet(), version.Version, port)
	})
	fmt.Println("Listening on :" + port)
	http.ListenAndServe(":"+port, nil)
}
