package main

import (
	// "fmt"
	"io"
	"net/http"
	"os"
)

/*
Get(), Post()，PostForm(), Head() all based on DefaultCLient()
Get() equals DefaultClient.Get()
*/
func main() {
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	req.Header.Add("User-Agent", "David User-Agent")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, resp.Body)
}
