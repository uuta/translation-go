package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const serverPort = 3333

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)

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
