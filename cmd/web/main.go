package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uuta/translation-go/internal/google"
)

// TODO: Does it need?
const serverPort = 3333

func main() {
	handleReq()

}

// handleReq is a function to make WEB server
func handleReq() {
	r := mux.NewRouter()
	// TODO: Fix endpoint
	// TODO: Make a handler
	r.HandleFunc("/", google.HandleGet())
	server := http.Server{
		// TODO: Consider serverPort is properly
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}
}
