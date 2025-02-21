package cmd

import "regexp"

// Regex pattern for URL extraction
var urlRegex = regexp.MustCompile(`https?://[^\s,]+`)

// Extract URLs from text
func ExtractURLs(text string) []string {
	return urlRegex.FindAllString(text, -1)
}
