package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	repo        = "dlawton/kuadrant-operator"
	baseURL     = "https://quay.io/api/v1/repository/"
)

var (
	robotPass = os.Getenv("ROBOT_PASS")
	robotUser = os.Getenv("ROBOT_USER")
)

type Tag struct {
	Name         string `json:"name"`
	LastModified string `json:"last_modified"`
}

type TagsResponse struct {
	Tags []Tag `json:"tags"`
}

func main() {
	client := &http.Client{}

	// Encode the robot credentials for Basic Auth
	auth := base64.StdEncoding.EncodeToString([]byte(robotUser + ":" + robotPass))
	
	// Create the request
	req, err := http.NewRequest("GET", baseURL+repo+"/tag", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Authorization", "Basic "+auth)

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Handle possible non-200 status codes
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\nBody: %s\n", resp.StatusCode, string(body))
		return
	}

	// Parse the JSON response
	var tagsResp TagsResponse
	if err := json.Unmarshal(body, &tagsResp); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	// Calculate the cutoff date for three weeks ago
	threeWeeksAgo := time.Now().AddDate(0, 0, -21)

	// Use a map to store unique tags (filter out duplicates)
	uniqueTags := make(map[string]struct{})
	for _, tag := range tagsResp.Tags {
		// Parse the LastModified timestamp
		lastModified, err := time.Parse(time.RFC1123, tag.LastModified)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			continue
		}

		// Filter out tags older than three weeks
		if lastModified.After(threeWeeksAgo) {
			uniqueTags[tag.Name] = struct{}{}
		}
	}

	// Print filtered tags
	for tag := range uniqueTags {
		fmt.Println(tag)
	}
}
