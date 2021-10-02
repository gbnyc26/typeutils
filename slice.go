package typeutils

import "strings"

func IndexStringSlice(slice []string) map[string]bool {
	// if slice is nil return an empty index

	r := make(map[string]bool)
	for _, s := range slice {
		r[s] = true
	}

	return r
}

func NilStringSliceEmpty(slice []string) []string {
	if slice == nil {
		return []string{}
	}

	return slice
}

func EmptyByteSliceNil(b []byte) []byte {
	if len(b) == 0 {
		return nil
	}
	return b
}

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
