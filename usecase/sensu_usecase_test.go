package usecase

import (
	"strings"
	"testing"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
	localtest "github.com/betorvs/sensubot/test"
	"github.com/stretchr/testify/assert"
)

func TestSensuURLGenerator(t *testing.T) {
	test1a := "check"
	test1b := "default"
	test1c := "all"
	result1 := sensuURLGenerator(test1a, test1b, test1c)
	if strings.Contains(result1, "namespaces/default/checks") {
		t.Fatalf("Invalid 10.1  TestSensuURLGenerator %s", result1)
	}

	test2a := "namespaces"
	test2b := "default"
	test2c := "all"
	result2 := sensuURLGenerator(test2a, test2b, test2c)
	assert.Contains(t, result2, "v2/namespaces")
	// if !strings.Contains(result2, "v2/namespaces") {
	// 	t.Fatalf("Invalid 11.1  TestSensuURLGenerator %s", result2)
	// }

	test3a := "silenced"
	test3b := "development"
	test3c := "all"
	result3 := sensuURLGenerator(test3a, test3b, test3c)
	if !strings.Contains(result3, "namespaces/development/silenced") {
		t.Fatalf("Invalid 12.1  TestSensuURLGenerator %s", result3)
	}

	test4a := "health"
	test4b := "default"
	test4c := "all"
	result4 := sensuURLGenerator(test4a, test4b, test4c)
	if strings.Contains(result4, "namespaces") {
		t.Fatalf("Invalid 13.1  TestSensuURLGenerator %s", result4)
	}
}

func TestRequestSensu(t *testing.T) {
	appcontext.Current.Add(appcontext.SensuRepository, localtest.InitSensuMock)
	appcontext.Current.Add(appcontext.Logger, localtest.InitMockLogger)
	command1 := make(map[string]string)
	command1["verb"] = "get"
	command1["name"] = "all"
	command1["resource"] = "checks"
	_ = requestSensu(command1, "xas111", "test user")
	expected1 := 0
	// expected 0 because variables are not configured
	assert.Equal(t, expected1, localtest.SensuCalls)
	// configuring variables to test the call
	config.Values.SensuAPI = "http://localhost:8080"
	config.Values.SensuAPIToken = "asdadasd"
	command2 := make(map[string]string)
	command2["verb"] = "get"
	command2["name"] = "all"
	command2["resource"] = "namespaces"
	_ = requestSensu(command2, "xas111", "test user")
	expected2 := 1
	assert.Equal(t, expected2, localtest.SensuCalls)

}

// func TestRequestSensuHealth(t *testing.T) {
// 	appcontext.Current.Add(appcontext.SensuRepository, localtest.InitSensuMock)
// 	// repo := SensuRepositoryMock{}
// 	// appcontext.Current.Add(appcontext.SensuRepository, repo)
// 	_ = localtest.SensuHealth("http://sensu-api:8080")
// 	expected := 1
// 	if sensuHealthRequestCalls != expected {
// 		t.Fatalf("Invalid 15.1 TestRequestSensuHealth %d", sensuHealthRequestCalls)
// 	}

// }
