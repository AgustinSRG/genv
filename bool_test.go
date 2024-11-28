// Tests

package genv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvBool(t *testing.T) {
	var val bool
	varName := "VAR_BOOL_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvBool(varName, true)
	assert.Equal(t, val, true)

	val = GetEnvBool(varName, false)
	assert.Equal(t, val, false)

	// Test TRUE

	os.Setenv(varName, "TRUE")

	val = GetEnvBool(varName, true)
	assert.Equal(t, val, true)

	val = GetEnvBool(varName, false)
	assert.Equal(t, val, true)

	// Test YES

	os.Setenv(varName, "YES")

	val = GetEnvBool(varName, true)
	assert.Equal(t, val, true)

	val = GetEnvBool(varName, false)
	assert.Equal(t, val, true)

	// Test FALSE

	os.Setenv(varName, "FALSE")

	val = GetEnvBool(varName, true)
	assert.Equal(t, val, false)

	val = GetEnvBool(varName, false)
	assert.Equal(t, val, false)

	// Test NO

	os.Setenv(varName, "NO")

	val = GetEnvBool(varName, true)
	assert.Equal(t, val, false)

	val = GetEnvBool(varName, false)
	assert.Equal(t, val, false)

	// Test case insensitive

	os.Setenv(varName, "yes")

	val = GetEnvBool(varName, false)
	assert.Equal(t, val, true)

	os.Setenv(varName, "no")

	val = GetEnvBool(varName, true)
	assert.Equal(t, val, false)

	// Test other

	os.Setenv(varName, "other")

	val = GetEnvBool(varName, false)
	assert.Equal(t, val, false)

	val = GetEnvBool(varName, true)
	assert.Equal(t, val, true)
}

func TestGetEnvBoolWithWarning(t *testing.T) {
	var val bool
	var warning bool
	varName := "VAR_BOOL_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvBoolWithWarning(varName, true)
	assert.Equal(t, val, true)
	assert.Equal(t, warning, false)

	val, warning = GetEnvBoolWithWarning(varName, false)
	assert.Equal(t, val, false)
	assert.Equal(t, warning, false)

	// Test TRUE

	os.Setenv(varName, "TRUE")

	val, warning = GetEnvBoolWithWarning(varName, true)
	assert.Equal(t, val, true)
	assert.Equal(t, warning, false)

	val, warning = GetEnvBoolWithWarning(varName, false)
	assert.Equal(t, val, true)
	assert.Equal(t, warning, false)

	// Test YES

	os.Setenv(varName, "YES")

	val, warning = GetEnvBoolWithWarning(varName, true)
	assert.Equal(t, val, true)
	assert.Equal(t, warning, false)

	val, warning = GetEnvBoolWithWarning(varName, false)
	assert.Equal(t, val, true)
	assert.Equal(t, warning, false)

	// Test FALSE

	os.Setenv(varName, "FALSE")

	val, warning = GetEnvBoolWithWarning(varName, true)
	assert.Equal(t, val, false)
	assert.Equal(t, warning, false)

	val, warning = GetEnvBoolWithWarning(varName, false)
	assert.Equal(t, val, false)
	assert.Equal(t, warning, false)

	// Test NO

	os.Setenv(varName, "NO")

	val, warning = GetEnvBoolWithWarning(varName, true)
	assert.Equal(t, val, false)
	assert.Equal(t, warning, false)

	val, warning = GetEnvBoolWithWarning(varName, false)
	assert.Equal(t, val, false)
	assert.Equal(t, warning, false)

	// Test case insensitive

	os.Setenv(varName, "yes")

	val, warning = GetEnvBoolWithWarning(varName, false)
	assert.Equal(t, val, true)
	assert.Equal(t, warning, false)

	os.Setenv(varName, "no")

	val, warning = GetEnvBoolWithWarning(varName, true)
	assert.Equal(t, val, false)
	assert.Equal(t, warning, false)

	// Test other

	os.Setenv(varName, "other")

	val, warning = GetEnvBoolWithWarning(varName, false)
	assert.Equal(t, val, false)
	assert.Equal(t, warning, true)

	val, warning = GetEnvBoolWithWarning(varName, true)
	assert.Equal(t, val, true)
	assert.Equal(t, warning, true)
}
