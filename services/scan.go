package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	mu sync.RWMutex
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

// ExtractURLs extracts URLs from the given text
func ExtractURLs(text string) []string {
	re := regexp.MustCompile(`https?://[^\s]+`)
	return re.FindAllString(text, -1)
}

// CheckURL checks if a URL is safe using Google Safe Browsing API
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

	matches, exists := result["matches"]
	fmt.Println("matches", matches)
	return !exists || len(matches.([]interface{})) == 0, nil
}

// ClassifyURLs checks URLs using Google Safe Browsing API
func ClassifyURLs(urls []string) map[string]string {
	mu.RLock()
	defer mu.RUnlock()

	classified := make(map[string]string)
	for _, url := range urls {
		isSafe, err := CheckURL(url)
		fmt.Println("IsSafe", isSafe)
		if err != nil {
			classified[url] = "error"
		} else if isSafe {
			classified[url] = "safe"
		} else {
			classified[url] = "suspicious"
		}
	}
	return classified
}

const (
	urlResultTemplate = `
%s URL Check: %s
→ Status: %s
→ Recommendation: %s
%s

`
)

func getRecommendationWithAction(classification string, url string) (string, string, string) {
	cleanURL := strings.Split(url, "\">")[0]

	switch classification {
	case "safe":
		action := fmt.Sprintf("→ Action: <a href=\"%s\" target=\"_blank\" rel=\"noopener noreferrer\" style=\"color: #0000EE; text-decoration: underline;\">Click here to visit website</a>", cleanURL)
		return "✅", "This link appears safe. You can proceed to visit it.", action
	case "suspicious":
		return "⚠️", "Exercise caution before visiting this link. It shows suspicious patterns.", "→ Action: URL hidden for your safety"
	case "error":
		return "❌", "Unable to verify this link's safety.", "→ Action: URL blocked - verification failed"
	default:
		return "❓", "Couldn't determine link safety.", "→ Action: URL hidden until verification"
	}
}

// func getRecommendationWithAction(classification string, url string) (string, string, string) {
// 	switch classification {
// 	case "safe":
// 		action := fmt.Sprintf("→ Action: Click here to visit: %s", url)
// 		return "✅", "This link appears safe. You can proceed to visit it.", action
// 	case "suspicious":
// 		return "⚠️", "Exercise caution before visiting this link. It shows suspicious patterns.", "→ Action: URL hidden for your safety"
// 	case "error":
// 		return "❌", "Unable to verify this link's safety.", "→ Action: URL blocked - verification failed"
// 	default:
// 		return "❓", "Couldn't determine link safety.", "→ Action: URL hidden until verification"
// 	}
// }

func BuildResponseMessage(msg string, classifications map[string]string) string {
	var sb strings.Builder

	for url, classification := range classifications {
		icon, recommendation, action := getRecommendationWithAction(classification, url)
		cleanURL := strings.Split(url, "\">")[0]
		displayURL := cleanURL
		if classification != "safe" {
			displayURL = "[URL Hidden]"
		}
		sb.WriteString(fmt.Sprintf(urlResultTemplate,
			icon,
			displayURL,
			classification,
			recommendation,
			action))
	}

	return sb.String()
}

// func BuildResponseMessage(msg string, classifications map[string]string) string {
// 	var sb strings.Builder

// 	for url, classification := range classifications {
// 		icon, recommendation := getRecommendation(classification)
// 		cleanURL := strings.Split(url, "\">")[0]
// 		sb.WriteString(fmt.Sprintf(urlResultTemplate,
// 			icon,
// 			cleanURL,
// 			classification,
// 			recommendation))
// 	}

// 	return sb.String()
// }

// func BuildResponseMessage(msg string, classifications map[string]string) string {
// 	var sb strings.Builder
// 	var safe, suspicious, errors int

// 	// Add header
// 	sb.WriteString(fmt.Sprintf(headerTemplate, msg))

// 	// Add URL analysis
// 	for url, classification := range classifications {
// 		icon := "✅"
// 		switch classification {
// 		case "suspicious":
// 			icon = "⚠️"
// 			suspicious++
// 		case "error":
// 			icon = "❌"
// 			errors++
// 		default:
// 			safe++
// 		}
// 		// Clean URL display by removing HTML artifacts
// 		cleanURL := strings.Split(url, "\">")[0]
// 		sb.WriteString(fmt.Sprintf(urlResultTemplate, icon, cleanURL, classification))
// 	}

// 	// Add summary
// 	total := len(classifications)
// 	sb.WriteString(fmt.Sprintf(summaryTemplate, total, safe, suspicious, errors))

// 	return sb.String()
// }
