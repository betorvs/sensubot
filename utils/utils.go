package utils

import "strings"

//StringInSlice checks if a slice contains a specific string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//Int64InSlice checks if a slice contains a specific int64
func Int64InSlice(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Trim func returns only the first n bytes of a string
func Trim(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

// ParseKeyValue receives an string to be splited and a string to be used to divide it
func ParseKeyValue(s, split string) (k, v string, r bool) {
	if strings.Contains(s, split) {
		t := strings.Split(s, split)
		if len(t) == 2 {
			k = t[0]
			v = t[1]
			r = true
			return k, v, r
		}
	}
	return "", "", false
}

// SearchLabels return a pair of strings, map[string]string and an boolean
func SearchLabels(key, value string, labels map[string]string) bool {
	if len(labels) == 0 {
		return false
	}
	for k, v := range labels {
		if k == key && v == value {
			return true
		}
	}
	return false
}

// ParseLabels receives a string divided by comma and return a map[string]string
func ParseLabels(s string) map[string]string {
	labels := map[string]string{}
	if strings.Contains(s, ",") {
		splittedString := strings.Split(s, ",")
		for _, v := range splittedString {
			key, value, valid := ParseKeyValue(v, "=")
			if valid {
				labels[key] = value
			}
		}
	}
	return labels
}
