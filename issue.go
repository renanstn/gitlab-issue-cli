package main

import (
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
)

const baseURL = "https://git.mop.equinix.com.br/api/v4"

type Issue struct {
	Title string `json:"title"`
	Url string `json:"web_url"`
	// Destription string `json:"description"`
}

func main() {
	// Check params
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <issue_number>")
		return
	}

	// Load environment vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		log.Fatal("Missing Gitlab token")
		return
	}

	// Concatenate issue prefix with number
	issueNumber := args[1]
	concatenatedString := "MSA-" + issueNumber
	searchURL := fmt.Sprintf("%s/issues?search=%s", baseURL, concatenatedString)

	// Get issue from gitlab
	request, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
		return
	}
	request.Header.Set("PRIVATE-TOKEN", token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Serialize response
	var issues []Issue
	err = json.NewDecoder(response.Body).Decode(&issues)
	if err != nil {
		log.Fatal("Error decoding response:", err)
		return
	}

	if len(issues) == 0 {
		fmt.Println("No matching issues found.")
		return
	}

	// Print results
	for _, issue := range issues {
		fmt.Println("- Title: " + issue.Title)
		fmt.Println("- URL:   " + issue.Url)
	}
}
