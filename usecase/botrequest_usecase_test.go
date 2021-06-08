package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	// verb specialWord|resource resouce|name specialWord namespace_name specialWord entity_name

	var validData1 = string("get checks")
	config1 := parseRequest(validData1)
	assert.Equal(t, "get", config1["verb"])
	assert.Equal(t, "checks", config1["resource"])

	var validData2 = string("get checks")
	config2 := parseRequest(validData2)
	assert.Equal(t, "all", config2["name"])
	assert.Equal(t, "default", config2["namespace"])

	var validData3 = string("get asset test")
	config3 := parseRequest(validData3)
	assert.Equal(t, "test", config3["name"])
	assert.Equal(t, "assets", config3["resource"])

	var validData4 = string("get assets test devnew")
	config4 := parseRequest(validData4)
	assert.Equal(t, "devnew", config4["namespace"])
	assert.Equal(t, "assets", config4["resource"])
	assert.Equal(t, "test", config4["name"])

	var validData5 = string("get asset test on stg")
	config5 := parseRequest(validData5)
	assert.Equal(t, "stg", config5["namespace"])
	assert.Equal(t, "test", config5["name"])

	var validData6 = string("silence check test on dev at server")
	config6 := parseRequest(validData6)
	assert.Equal(t, "dev", config6["namespace"])
	assert.Equal(t, "silence", config6["verb"])
	assert.Equal(t, "server", config6["entity"])

	// Test with strange string behaviours
	var test1 = string("get CHecks")
	testConfig1 := parseRequest(test1)
	assert.Equal(t, "get", testConfig1["verb"])
	assert.Equal(t, "all", testConfig1["name"])

	var test2 = string("Get silences")
	testConfig2 := parseRequest(test2)
	assert.Equal(t, "all", testConfig2["name"])
	assert.Equal(t, "silenced", testConfig2["resource"])

	var test3 = string("Get Namespaces")
	testConfig3 := parseRequest(test3)
	assert.Equal(t, "all", testConfig3["name"])
	assert.Equal(t, "namespaces", testConfig3["resource"])
	assert.Equal(t, "get", testConfig3["verb"])

	test4 := "get alertmanager.alerts namespace=argocd"
	testConfig4 := parseRequest(test4)
	assert.Equal(t, "get", testConfig4["verb"])
	assert.Equal(t, "alertmanager", testConfig4["integration"])

	test5 := "get default.events namespace=argocd sensu-alertmanager-events=owner"
	testConfig5 := parseRequest(test5)
	assert.Equal(t, "get", testConfig5["verb"])
	assert.Equal(t, "sensu", testConfig5["integration"])
	assert.Equal(t, "namespace=argocd,sensu-alertmanager-events=owner", testConfig5["labels"])
}
