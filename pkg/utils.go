package pkg

import "strings"

const vscSeparator = ","

func someHelperFunction(s string) []string {
	result := strings.Split(s, vscSeparator)
	for i, part := range result {
		result[i] = strings.ReplaceAll(part, ";", "")
	}
	return result
}
