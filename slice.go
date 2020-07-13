package typeutils

import "strings"

func StringSplitEmptySlice(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

func StringSplitNilSlice(s string, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}
