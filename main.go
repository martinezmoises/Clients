// Filename: main.go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//to where we are making the request(client)
	url := "https://api.github.com" //API github.com (Server)

	//creat a request(new request)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	//creating a client
	client := http.Client{} //This does not know how long it would wait for response
	//Generic client        //Type called type creating an instances of this type called client
	response, err := client.Do(request) //This is sending request (API GITHUB.COM)
	if err != nil {
		panic(err) //Program stops
	}
	//we need to close the response
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body) //Returns bytes
	if err != nil {
		panic(err)
	}

	//IF this codes executes we did not get an error (Response is valid)
	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes)) //Converts bytes so it can be readable
}
