package typeutils

import "strings"

func SplitEmptyStringEmptySlice(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

func SplitEmptyStringNilSlice(s string, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}
