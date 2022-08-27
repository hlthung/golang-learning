package slicehelper

import (
	"strings"
)

func RemoveDuplicateStrings(in []string) (out []string, dupes []string) {
	keys := make(map[string]bool)
	for _, x := range in {
		if _, ok := keys[x]; !ok {
			keys[x] = true
			out = append(out, x)
		} else {
			dupes = append(dupes, x)
		}
	}
	return out, dupes
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsStringInsensitive(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}

func FilterStrings(in []string, f func(s string) bool) []string {
	var out []string
	for _, s := range in {
		if !f(s) {
			out = append(out, s)
		}
	}
	return out
}

func MapStrings(slice []string, f func(str string) string) []string {
	var out = make([]string, len(slice))
	for i := range slice {
		out[i] = f(slice[i])
	}
	return out
}

func Equals(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for idx := range s1 {
		if s1[idx] != s2[idx] {
			return false
		}
	}
	return true
}
