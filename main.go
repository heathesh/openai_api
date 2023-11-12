package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter prompt: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You entered: %s-\n", scanner.Text())

	settings := readSettings()

	maxTokensStr := settings["completionMaxTokens"]
	maxTokens, err := strconv.ParseInt(maxTokensStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a request body
	requestBody := CompletionAPIRequest{
		Model:     settings["completionModel"],
		Messages:  []CompletionAPIRequestMessage{{Role: "user", Content: []CompletionAPIRequestContent{{Type: "text", Text: scanner.Text()}}}},
		MaxTokens: maxTokens,
	}

	// call the Marshal method on a pointer to the instance
	data, err := (&requestBody).Marshal()
	if err != nil {
		log.Fatal(err)
	}

	// create a new HTTP POST request
	req, err := http.NewRequest("POST", settings["completionsEndpoint"], bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	// set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+settings["apiKey"])

	// create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := UnmarshalCompletionAPIResponse(body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// save the response in the output.md file
	fileWriteErr := os.WriteFile("output.md", []byte(response.Choices[0].Message.Content), 0644)
	if fileWriteErr != nil {
		log.Fatal(fileWriteErr)
	}

	// output the response body
	fmt.Println(response.Choices[0].Message.Content)
}

func readSettings() map[string]string {
	// read the settings file
	data, err := os.ReadFile("settings.json")
	if err != nil {
		log.Fatal(err)
	}

	// parse the JSON data
	var settings map[string]string
	err = json.Unmarshal(data, &settings)
	if err != nil {
		log.Fatal(err)
	}

	return settings
}
