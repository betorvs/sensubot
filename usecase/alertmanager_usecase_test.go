package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestAlertManager(t *testing.T) {
	command := map[string]string{}
	res1 := requestAlertManager(command, "admin", "Test Admin")
	assert.Equal(t, "SensuBot Not Configured to access any Alert Manager Endpoint", res1)
}
