package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// resp, err := http.Get("https://httpbin.org/get")
	// resp, err := http.Get("https://httpbin.org/get?search=something")
	resp, err := http.Post("https://httpbin.org/post", "text/plain",
		strings.NewReader("this is thhe request content"))
	if err != nil {
		log.Fatalln("Unable to get response")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}
	fmt.Println(string(content))
}
