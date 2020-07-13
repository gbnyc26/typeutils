package typeutils

import "strings"

func SplitEmptyStringSliceEmpty(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

func SplitEmptyStringSliceNil(s string, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}
