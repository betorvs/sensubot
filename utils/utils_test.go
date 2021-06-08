package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	testSlice := []string{"foo", "bar", "test"}
	testString := "test"
	testResult := StringInSlice(testString, testSlice)
	assert.True(t, testResult)
	testSlice1 := []string{"foo", "bar"}
	testResult1 := StringInSlice(testString, testSlice1)
	assert.False(t, testResult1)

}

func TestInt64InSlice(t *testing.T) {
	testSlice := []int64{123123123, 123123123133}
	testInt64 := int64(123123123133)
	testResult := Int64InSlice(testInt64, testSlice)
	assert.True(t, testResult)
	test2Int64 := int64(123123133)
	testResult2 := Int64InSlice(test2Int64, testSlice)
	assert.False(t, testResult2)
}

func TestTrim(t *testing.T) {
	test1 := "abcde"
	res1 := Trim(test1, 2)
	assert.Equal(t, "ab", res1)
	test2 := "ab"
	res2 := Trim(test2, 2)
	assert.Equal(t, "ab", res2)
}

func TestParseKeyValue(t *testing.T) {
	test1 := "foo:bar"
	_, _, res1 := ParseKeyValue(test1, "=")
	assert.False(t, res1)
	_, _, res2 := ParseKeyValue(test1, ":")
	assert.True(t, res2)
}

func TestSearchLabels(t *testing.T) {
	labels1 := map[string]string{}
	res1 := SearchLabels("foo", "bar", labels1)
	assert.False(t, res1)
	labels2 := map[string]string{"foo": "bar", "local": "host"}
	res2 := SearchLabels("foo", "bar", labels2)
	assert.True(t, res2)
	res3 := SearchLabels("foo", "host", labels2)
	assert.False(t, res3)
}

func TestParseLabels(t *testing.T) {
	test1 := "foo=bar,local=host"
	labels1 := map[string]string{"foo": "bar", "local": "host"}
	res1 := ParseLabels(test1)
	assert.Equal(t, labels1, res1)

}
