package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestOpsgenie(t *testing.T) {
	command := map[string]string{}
	res1 := requestOpsgenie(command, "admin", "Test Admin")
	assert.Equal(t, "SensuBot Not Configured to access Opsgenie", res1)
}
