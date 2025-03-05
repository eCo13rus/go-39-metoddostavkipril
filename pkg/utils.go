package pkg

import "strings"

const vscSeparator = ","

func someHelperFunction(s string) []string {
	s = strings.ReplaceAll(s, ";", ",")
	return strings.Split(s, vscSeparator)
}
