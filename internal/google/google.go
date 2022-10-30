package google

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (h *handler) HandleGet(w http.ResponseWriter, r *http.Request) error {
	// TODO: make a function that deals with downloading
	// Make a file
	out, err := os.Create("translation.mp3")
	if err != nil {
		return fmt.Errorf("error os create: %s\n", err)
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
	return nil
}
