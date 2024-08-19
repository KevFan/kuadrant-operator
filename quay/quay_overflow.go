package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	repo    = "test_org_123/kuadrant-operator"
	baseURL = "https://quay.io/api/v1/repository/"
)

var (
	robotPass   = os.Getenv("ROBOT_PASS")
	robotUser   = os.Getenv("ROBOT_USER")
	accessToken = os.Getenv("ACCESS_TOKEN")
	preserveSubstring = "danlaw345"
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

	// Create the request to get tags
	req, err := http.NewRequest("GET", baseURL+repo+"/tag", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Prioritize Bearer token for authorization
	if accessToken != "" {
		req.Header.Add("Authorization", "Bearer "+accessToken)
	} else {
		// Fallback to Basic Authentication if no access token
		auth := base64.StdEncoding.EncodeToString([]byte(robotUser + ":" + robotPass))
		req.Header.Add("Authorization", "Basic "+auth)
	}

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
	cutOffTime := time.Now().AddDate(0, 1, 0)

	// Use a map to store unique tags (filter out duplicates)
	uniqueTags := make(map[string]struct{})
	for _, tag := range tagsResp.Tags {
		// Parse the LastModified timestamp
		lastModified, err := time.Parse(time.RFC1123, tag.LastModified)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			continue
		}

	// Delete tags older than three hours, unless they contain the preserve substring
	if lastModified.Before(cutOffTime) && !containsSubstring(tag.Name, preserveSubstring) {
		deleteTag(client, accessToken, tag.Name)
	} else {
		uniqueTags[tag.Name] = struct{}{}
	}
	}

	// Print remaining tags
	for tag := range uniqueTags {
		fmt.Println(tag)
	}
}

func containsSubstring(tagName, substring string) bool {
	return strings.Contains(tagName, substring)
}

// deleteTag sends a DELETE request to remove the specified tag from the repository
func deleteTag(client *http.Client, accessToken, tagName string) {
	req, err := http.NewRequest("DELETE", baseURL+repo+"/tag/"+tagName, nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error deleting tag:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		fmt.Printf("Successfully deleted tag: %s\n", tagName)
	} else {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Failed to delete tag %s: Status code %d\nBody: %s\n", tagName, resp.StatusCode, string(body))
	}
	
}
