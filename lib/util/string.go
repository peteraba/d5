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

func JoinLimited(parts []string, maxCount int) string {
	if len(parts) < maxCount {
		maxCount = len(parts)
	}

	if maxCount < 1 {
		return ""
	}

	return strings.Join(parts[0:maxCount], ", ")
}
