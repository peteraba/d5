package util

import "strings"

func TrimSplit(s, sep string) []string {
	if strings.Trim(s, " \n\t") == "" {
		return []string{}
	}

	split := strings.Split(s, sep)

	for key, word := range split {
		split[key] = strings.Trim(word, " \n\t")
	}

	return split
}

func JoinLimited(parts []string, joinBy string, maxCount int) string {
	if len(parts) < maxCount {
		maxCount = len(parts)
	}

	if maxCount < 1 {
		return ""
	}

	return strings.Join(parts[0:maxCount], joinBy)
}

func HasSuffixAny(s string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}

	return false
}

func StringIn(s string, options []string) bool {
	for _, option := range options {
		if option == s {
			return true
		}
	}

	return false
}
