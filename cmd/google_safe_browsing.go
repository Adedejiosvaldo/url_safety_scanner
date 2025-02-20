package cmd

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

const (
	safeBrowsingAPIURL = "https://safebrowsing.googleapis.com/v4/threatMatches:find"
)

type ThreatRequest struct {
	Client struct {
		ClientID      string `json:"clientId"`
		ClientVersion string `json:"clientVersion"`
	} `json:"client"`
	ThreatInfo struct {
		ThreatTypes      []string `json:"threatTypes"`
		PlatformTypes    []string `json:"platformTypes"`
		ThreatEntryTypes []string `json:"threatEntryTypes"`
		ThreatEntries    []struct {
			URL string `json:"url"`
		} `json:"threatEntries"`
	} `json:"threatInfo"`
}

func CheckURL(urlToCheck string) (bool, error) {

	API_KEY := os.Getenv("GOOGLE_SAFE_BROWSING_API_KEY")

	request := ThreatRequest{}
	request.Client.ClientID = "your-client-name"
	request.Client.ClientVersion = "1.0.0"

	request.ThreatInfo.ThreatTypes = []string{
		"MALWARE",
		"SOCIAL_ENGINEERING",
		"UNWANTED_SOFTWARE",
		"POTENTIALLY_HARMFUL_APPLICATION",
	}
	request.ThreatInfo.PlatformTypes = []string{"ANY_PLATFORM"}
	request.ThreatInfo.ThreatEntryTypes = []string{"URL"}
	request.ThreatInfo.ThreatEntries = []struct {
		URL string `json:"url"`
	}{{URL: urlToCheck}}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", safeBrowsingAPIURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return false, err
	}

	q := req.URL.Query()
	q.Add("key", API_KEY)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	// If matches is empty, the URL is safe
	matches, exists := result["matches"]
	return !exists || len(matches.([]interface{})) == 0, nil
}
