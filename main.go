// Filename: main.go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//to where we are making the request(client)
	url := "https://api.github.com"

	//creating a client
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	//we need to close the response
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}
