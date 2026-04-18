package utils

import "regexp"

var htmlTags = regexp.MustCompile(`<[^>]*>`)

// StripHTML removes HTML tags from a string, leaving plain text.
func StripHTML(s string) string {
	return htmlTags.ReplaceAllString(s, "")
}
