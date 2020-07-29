package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// way 1 to get data
	data := make([]byte, 999999)
	fmt.Println(resp.Body.Read(data))
	fmt.Println(string(data))

	// way 2 to get data
	io.Copy(os.Stdout, resp.Body)

	// way 3 to get data
	lw := logWriter{}
	// lw implements wtire so it's ok to use it
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
