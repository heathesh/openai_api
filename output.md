To call a REST API in Golang, you can use the `net/http` package which provides HTTP client and server implementations. Below is an example code to call a REST API using the `GET` method:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// URL of the REST API
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating the HTTP request:", err)
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Perform the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error performing the HTTP request:", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading the response body:", err)
	}

	// Print the response body
	fmt.Println(string(body))
}
```

This code sends a `GET` request to the specified URL and reads the response body. It then prints the response body to the console. Remember to handle errors properly in a real-world scenario and close the response body using `defer resp.Body.Close()` to prevent resource leaks.