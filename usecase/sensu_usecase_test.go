package usecase

import (
	"strings"
	"testing"

	"github.com/betorvs/sensubot/appcontext"
)

var (
	sensuGetRequestCalls    int
	sensuPostRequestCalls   int
	sensuHealthRequestCalls int
)

func TestParseRequest(t *testing.T) {
	// verb specialWord|resource resouce|name specialWord namespace_name specialWord entity_name

	var validData1 = string("get checks")
	config1 := parseRequest(validData1)
	if config1["verb"] != "get" {
		t.Fatalf("Invalid 1.1 verb associate %s", config1["verb"])
	}
	if config1["resource"] != "checks" {
		t.Fatalf("Invalid 1.2 verb associate %s", config1["resource"])
	}

	var validData2 = string("get all checks")
	config2 := parseRequest(validData2)
	if config2["name"] != "all" {
		t.Fatalf("Invalid 2.1 name associate %s", config2["name"])
	}
	if config2["namespace"] != "default" {
		t.Fatalf("Invalid 2.2 namespace associate %s", config2["namespace"])
	}

	var validData3 = string("get asset test")
	config3 := parseRequest(validData3)
	if config3["name"] != "test" {
		t.Fatalf("Invalid 3.1 name associate %s", config3["name"])
	}
	if config3["resource"] != "assets" {
		t.Fatalf("Invalid 3.2 resource associate %s", config3["resource"])
	}

	var validData4 = string("get assets test devnew")
	config4 := parseRequest(validData4)
	if config4["namespace"] != "devnew" {
		t.Fatalf("Invalid 4.1 namespace associate %s", config4["namespace"])
	}
	if config4["resource"] != "assets" {
		t.Fatalf("Invalid 4.2 assets associate %s", config4["resource"])
	}
	if config4["name"] != "test" {
		t.Fatalf("Invalid 4.3 name associate %s", config4["name"])
	}

	var validData5 = string("get asset test on stg")
	config5 := parseRequest(validData5)
	if config5["namespace"] != "stg" {
		t.Fatalf("Invalid 5.1 namespace associate %s", config5["namespace"])
	}
	if config5["name"] != "test" {
		t.Fatalf("Invalid 5.2 name associate %s", config5["name"])
	}

	var validData6 = string("silence check test on dev at server")
	config6 := parseRequest(validData6)
	if config6["namespace"] != "dev" {
		t.Fatalf("Invalid 6.1 namespace associate %s", config6["namespace"])
	}
	if config6["verb"] != "silence" {
		t.Fatalf("Invalid 6.2 verb associate %s", config6["verb"])
	}
	if config6["entity"] != "server" {
		t.Fatalf("Invalid 6.3 entity associate %s", config6["entity"])
	}

	// Test with strange string behaviours
	var test1 = string(" get  all CHecks")
	testConfig1 := parseRequest(test1)
	if testConfig1["verb"] != "get" {
		t.Fatalf("Invalid 7.1 verb associate %s", testConfig1["verb"])
	}
	if testConfig1["name"] != "all" {
		t.Fatalf("Invalid 7.2 name associate %s", testConfig1["name"])
	}

	var test2 = string("Get   all   silences")
	testConfig2 := parseRequest(test2)
	if testConfig2["name"] != "all" {
		t.Fatalf("Invalid 8.1 name associate %s", testConfig2["name"])
	}
	if testConfig2["resource"] != "silenced" {
		t.Fatalf("Invalid 8.2 resouces associate %s", testConfig2["resource"])
	}

	var test3 = string("Get   all   Namespaces")
	testConfig3 := parseRequest(test3)
	if testConfig3["name"] != "all" {
		t.Fatalf("Invalid 9.1 name associate %s", testConfig3["name"])
	}
	if testConfig3["resource"] != "namespaces" {
		t.Fatalf("Invalid 9.2 resouces associate %s", testConfig3["resource"])
	}
}

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
	if !strings.Contains(result2, "v2/namespaces") {
		t.Fatalf("Invalid 11.1  TestSensuURLGenerator %s", result2)
	}

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

type SensuRepositoryMock struct {
}

func (repo SensuRepositoryMock) SensuGet(sensuurl string, token string) ([]byte, error) {
	sensuGetRequestCalls++
	return []byte{}, nil
}

func (repo SensuRepositoryMock) SensuPost(sensuurl string, token string, body []byte) ([]byte, error) {
	sensuPostRequestCalls++
	return []byte{}, nil
}

func (repo SensuRepositoryMock) SensuHealth(sensuurl string) bool {
	sensuHealthRequestCalls++
	return true
}

func TestRequestSensu(t *testing.T) {
	repo := SensuRepositoryMock{}
	appcontext.Current.Add(appcontext.SensuRepository, repo)
	_ = requestSensu("get all namespaces")
	expected := 1
	if sensuGetRequestCalls != expected {
		t.Fatalf("Invalid 14.1 TestRequestSensu %d", sensuGetRequestCalls)
	}

}

func TestRequestSensuHealth(t *testing.T) {
	repo := SensuRepositoryMock{}
	appcontext.Current.Add(appcontext.SensuRepository, repo)
	_ = repo.SensuHealth("http://sensu-api:8080")
	expected := 1
	if sensuHealthRequestCalls != expected {
		t.Fatalf("Invalid 15.1 TestRequestSensuHealth %d", sensuHealthRequestCalls)
	}

}
