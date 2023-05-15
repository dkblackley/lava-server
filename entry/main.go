package main

import (
	"errors"
	"fmt"
	"lava_server/http_server"
	"net/http"
	"os"
)

// The nonce for chacha should be a running hash of the original mp4 video and the previous key.
func main() {
	http.HandleFunc("/", http_server.Get_root)
	http.HandleFunc("/hello", http_server.Get_hello)

	var err error
	err = http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
