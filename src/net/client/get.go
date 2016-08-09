package main

import (
	// "fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
