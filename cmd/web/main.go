package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// TODO: Does it need?
const serverPort = 3333

func main() {
	handleReq()

	time.Sleep(100 * time.Millisecond)

	// TODO: make a function that deals with downloading
	// Make a file
	out, err := os.Create("translation.mp3")
	if err != nil {
		fmt.Printf("error os create: %s\n", err)
	}
	defer out.Close()

	url := "https://translate.google.com/translate_tts?ie=UTF-8&client=tw-ob&q=Understand&tl=en&total=1&idx=0"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "audio/mpeg")
	req.Header.Set("Content-Disposition", "attachment; filename='translation.mp3'")
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	io.Copy(out, res.Body)
}

// handleReq is a function to make WEB server
func handleReq() {
	r := mux.NewRouter()
	// TODO: Fix endpoint
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
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
