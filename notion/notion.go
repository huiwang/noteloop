package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func UploadPage() {
	// Define the Notion API endpoint
	apiURL := "https://api.notion.com/v1/pages"

	// Define the request payload
	payload := map[string]interface{}{
		"parent": map[string]interface{}{
			"type":        "database_id",
			"database_id": "",
		},
		"properties": map[string]interface{}{
			"Note": map[string]interface{}{
				"title": []map[string]interface{}{
					{
						"text": map[string]interface{}{
							"content": "New Page Title",
						},
					},
				},
			},
		},
	}

	// Convert the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer ")
	req.Header.Set("Notion-Version", "2022-06-28")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error while reading books body: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		fmt.Println("body: ", string(body))
		return
	}

	fmt.Println("Page inserted successfully!")
}
