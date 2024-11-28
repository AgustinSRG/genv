// Tests

package genv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvFloat32(t *testing.T) {
	var val float32
	varName := "VAR_FLOAT32_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvFloat32(varName, 0.0)
	assert.Equal(t, val, float32(0.0))

	val = GetEnvFloat32(varName, 1.5)
	assert.Equal(t, val, float32(1.5))

	// Test valid value

	os.Setenv(varName, "1.2345")

	val = GetEnvFloat32(varName, 0.0)
	assert.Equal(t, val, float32(1.2345))

	val = GetEnvFloat32(varName, 1.5)
	assert.Equal(t, val, float32(1.2345))

	// Test invalid value

	os.Setenv(varName, "abcd")

	val = GetEnvFloat32(varName, 0.0)
	assert.Equal(t, val, float32(0.0))

	val = GetEnvFloat32(varName, 1.5)
	assert.Equal(t, val, float32(1.5))

	// Test with min value

	os.Setenv(varName, "-3.4e+38")

	val = GetEnvFloat32(varName, 0.0)
	assert.Equal(t, val, float32(-3.4e+38))

	val = GetEnvFloat32(varName, 1.5)
	assert.Equal(t, val, float32(-3.4e+38))

	// Test with max value

	os.Setenv(varName, "3.4e+38")

	val = GetEnvFloat32(varName, 0.0)
	assert.Equal(t, val, float32(3.4e+38))

	val = GetEnvFloat32(varName, 1.5)
	assert.Equal(t, val, float32(3.4e+38))

	// Test with overflow value

	os.Setenv(varName, "3.4e+39")

	val = GetEnvFloat32(varName, 0.0)
	assert.Equal(t, val, float32(0.0))

	val = GetEnvFloat32(varName, 1.5)
	assert.Equal(t, val, float32(1.5))
}

func TestGetEnvFloat32WithWarning(t *testing.T) {
	var val float32
	var warning bool
	varName := "VAR_FLOAT32_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvFloat32WithWarning(varName, 0.0)
	assert.Equal(t, val, float32(0.0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvFloat32WithWarning(varName, 1.5)
	assert.Equal(t, val, float32(1.5))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "1.2345")

	val, warning = GetEnvFloat32WithWarning(varName, 0.0)
	assert.Equal(t, val, float32(1.2345))
	assert.Equal(t, warning, false)

	val, warning = GetEnvFloat32WithWarning(varName, 1.5)
	assert.Equal(t, val, float32(1.2345))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "abcd")

	val, warning = GetEnvFloat32WithWarning(varName, 0.0)
	assert.Equal(t, val, float32(0.0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvFloat32WithWarning(varName, 1.5)
	assert.Equal(t, val, float32(1.5))
	assert.Equal(t, warning, true)

	// Test with min value

	os.Setenv(varName, "-3.4e+38")

	val, warning = GetEnvFloat32WithWarning(varName, 0.0)
	assert.Equal(t, val, float32(-3.4e+38))
	assert.Equal(t, warning, false)

	val, warning = GetEnvFloat32WithWarning(varName, 1.5)
	assert.Equal(t, val, float32(-3.4e+38))
	assert.Equal(t, warning, false)

	// Test with max value

	os.Setenv(varName, "3.4e+38")

	val, warning = GetEnvFloat32WithWarning(varName, 0.0)
	assert.Equal(t, val, float32(3.4e+38))
	assert.Equal(t, warning, false)

	val, warning = GetEnvFloat32WithWarning(varName, 1.5)
	assert.Equal(t, val, float32(3.4e+38))
	assert.Equal(t, warning, false)

	// Test with overflow value

	os.Setenv(varName, "3.4e+39")

	val, warning = GetEnvFloat32WithWarning(varName, 0.0)
	assert.Equal(t, val, float32(0.0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvFloat32WithWarning(varName, 1.5)
	assert.Equal(t, val, float32(1.5))
	assert.Equal(t, warning, true)
}

func TestGetEnvFloat64(t *testing.T) {
	var val float64
	varName := "VAR_FLOAT64_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvFloat64(varName, 0.0)
	assert.Equal(t, val, float64(0.0))

	val = GetEnvFloat64(varName, 1.5)
	assert.Equal(t, val, float64(1.5))

	// Test valid value

	os.Setenv(varName, "1.2345")

	val = GetEnvFloat64(varName, 0.0)
	assert.Equal(t, val, float64(1.2345))

	val = GetEnvFloat64(varName, 1.5)
	assert.Equal(t, val, float64(1.2345))

	// Test invalid value

	os.Setenv(varName, "abcd")

	val = GetEnvFloat64(varName, 0.0)
	assert.Equal(t, val, float64(0.0))

	val = GetEnvFloat64(varName, 1.5)
	assert.Equal(t, val, float64(1.5))

	// Test with min value

	os.Setenv(varName, "-1.7e+308")

	val = GetEnvFloat64(varName, 0.0)
	assert.Equal(t, val, float64(-1.7e+308))

	val = GetEnvFloat64(varName, 1.5)
	assert.Equal(t, val, float64(-1.7e+308))

	// Test with max value

	os.Setenv(varName, "1.7e+308")

	val = GetEnvFloat64(varName, 0.0)
	assert.Equal(t, val, float64(1.7e+308))

	val = GetEnvFloat64(varName, 1.5)
	assert.Equal(t, val, float64(1.7e+308))

	// Test with overflow value

	os.Setenv(varName, "1.7e+309")

	val = GetEnvFloat64(varName, 0.0)
	assert.Equal(t, val, float64(0.0))

	val = GetEnvFloat64(varName, 1.5)
	assert.Equal(t, val, float64(1.5))
}

func TestGetEnvFloat64WithWarning(t *testing.T) {
	var val float64
	var warning bool
	varName := "VAR_FLOAT64_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvFloat64WithWarning(varName, 0.0)
	assert.Equal(t, val, float64(0.0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvFloat64WithWarning(varName, 1.5)
	assert.Equal(t, val, float64(1.5))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "1.2345")

	val, warning = GetEnvFloat64WithWarning(varName, 0.0)
	assert.Equal(t, val, float64(1.2345))
	assert.Equal(t, warning, false)

	val, warning = GetEnvFloat64WithWarning(varName, 1.5)
	assert.Equal(t, val, float64(1.2345))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "abcd")

	val, warning = GetEnvFloat64WithWarning(varName, 0.0)
	assert.Equal(t, val, float64(0.0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvFloat64WithWarning(varName, 1.5)
	assert.Equal(t, val, float64(1.5))
	assert.Equal(t, warning, true)

	// Test with min value

	os.Setenv(varName, "-1.7e+308")

	val, warning = GetEnvFloat64WithWarning(varName, 0.0)
	assert.Equal(t, val, float64(-1.7e+308))
	assert.Equal(t, warning, false)

	val, warning = GetEnvFloat64WithWarning(varName, 1.5)
	assert.Equal(t, val, float64(-1.7e+308))
	assert.Equal(t, warning, false)

	// Test with max value

	os.Setenv(varName, "1.7e+308")

	val, warning = GetEnvFloat64WithWarning(varName, 0.0)
	assert.Equal(t, val, float64(1.7e+308))
	assert.Equal(t, warning, false)

	val, warning = GetEnvFloat64WithWarning(varName, 1.5)
	assert.Equal(t, val, float64(1.7e+308))
	assert.Equal(t, warning, false)

	// Test with overflow value

	os.Setenv(varName, "1.7e+309")

	val, warning = GetEnvFloat64WithWarning(varName, 0.0)
	assert.Equal(t, val, float64(0.0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvFloat64WithWarning(varName, 1.5)
	assert.Equal(t, val, float64(1.5))
	assert.Equal(t, warning, true)
}
