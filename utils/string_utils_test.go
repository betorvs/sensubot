package utils

import "testing"

func TestStringInSlice(t *testing.T) {
	testSlice := []string{"foo", "bar", "test"}
	testString := "test"
	testResult := StringInSlice(testString, testSlice)
	if !testResult {
		t.Fatalf("Invalid 1.1 TestStringInSlice %v", testResult)
	}

}
