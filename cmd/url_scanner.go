package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func UseURLScannerAPI(url string) (map[string]interface{}, error) {

	API_KEY := os.Getenv("URL_SCAN_KEY")
	postBody, _ := json.Marshal(map[string]string{
		"url":        url,
		"visibility": "public",
	})

	responseBody := bytes.NewBuffer(postBody)

	res, _ := http.NewRequest(
		http.MethodPost,
		"https://urlscan.io/api/v1/scan/", responseBody)
	res.Header.Add("API-Key", API_KEY)
	res.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(res)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	// cleanJSON, err := json.Marshal(result)

	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}
