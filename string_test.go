// Tests

package genv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvString(t *testing.T) {
	var val string
	varName := "VAR_STRING_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvString(varName, "")
	assert.Equal(t, val, "")

	val = GetEnvString(varName, "default")
	assert.Equal(t, val, "")

	// Test not set

	os.Unsetenv(varName)

	val = GetEnvString(varName, "")
	assert.Equal(t, val, "")

	val = GetEnvString(varName, "default")
	assert.Equal(t, val, "default")

	// Test with value

	os.Setenv(varName, "value")

	val = GetEnvString(varName, "")
	assert.Equal(t, val, "value")

	val = GetEnvString(varName, "default")
	assert.Equal(t, val, "value")
}

func TestGetEnvStringNonEmpty(t *testing.T) {
	var val string
	varName := "VAR_STRING_2"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvStringNonEmpty(varName, "")
	assert.Equal(t, val, "")

	val = GetEnvStringNonEmpty(varName, "default")
	assert.Equal(t, val, "default")

	// Test not set

	os.Unsetenv(varName)

	val = GetEnvStringNonEmpty(varName, "")
	assert.Equal(t, val, "")

	val = GetEnvStringNonEmpty(varName, "default")
	assert.Equal(t, val, "default")

	// Test with value

	os.Setenv(varName, "value")

	val = GetEnvStringNonEmpty(varName, "")
	assert.Equal(t, val, "value")

	val = GetEnvStringNonEmpty(varName, "default")
	assert.Equal(t, val, "value")
}
