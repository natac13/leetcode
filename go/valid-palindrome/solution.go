package validpalindrome

import (
	"log"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	s = removeNonAlphanumeric(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			log.Printf("s[i] = %v, s[len(s)-1-i] = %v", s[i], s[len(s)-1-i])
			return false
		}
	}

	return true
}

func removeNonAlphanumeric(s string) string {
	re := regexp.MustCompile(`[^a-z0-9]`)
	return re.ReplaceAllString(s, "")
}
