// Tests

package genv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvInt(t *testing.T) {
	var val int
	varName := "VAR_INT_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvInt(varName, 0)
	assert.Equal(t, val, int(0))

	val = GetEnvInt(varName, 2)
	assert.Equal(t, val, int(2))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvInt(varName, 0)
	assert.Equal(t, val, int(5))

	val = GetEnvInt(varName, 2)
	assert.Equal(t, val, int(5))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvInt(varName, 0)
	assert.Equal(t, val, int(0))

	val = GetEnvInt(varName, 2)
	assert.Equal(t, val, int(0))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvInt(varName, 0)
	assert.Equal(t, val, int(0))

	val = GetEnvInt(varName, 2)
	assert.Equal(t, val, int(2))
}

func TestGetEnvIntWithWarning(t *testing.T) {
	var val int
	var warning bool
	varName := "VAR_INT_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvIntWithWarning(varName, 0)
	assert.Equal(t, val, int(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvIntWithWarning(varName, 2)
	assert.Equal(t, val, int(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvIntWithWarning(varName, 0)
	assert.Equal(t, val, int(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvIntWithWarning(varName, 2)
	assert.Equal(t, val, int(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvIntWithWarning(varName, 0)
	assert.Equal(t, val, int(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvIntWithWarning(varName, 2)
	assert.Equal(t, val, int(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvIntWithWarning(varName, 0)
	assert.Equal(t, val, int(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvIntWithWarning(varName, 2)
	assert.Equal(t, val, int(2))
	assert.Equal(t, warning, true)
}

func TestGetEnvInt8(t *testing.T) {
	var val int8
	varName := "VAR_INT8_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvInt8(varName, 0)
	assert.Equal(t, val, int8(0))

	val = GetEnvInt8(varName, 2)
	assert.Equal(t, val, int8(2))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvInt8(varName, 0)
	assert.Equal(t, val, int8(0))

	val = GetEnvInt8(varName, 2)
	assert.Equal(t, val, int8(0))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvInt8(varName, 0)
	assert.Equal(t, val, int8(5))

	val = GetEnvInt8(varName, 2)
	assert.Equal(t, val, int8(5))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvInt8(varName, 0)
	assert.Equal(t, val, int8(0))

	val = GetEnvInt8(varName, 2)
	assert.Equal(t, val, int8(2))

	// Test with min value

	os.Setenv(varName, "-128")

	val = GetEnvInt8(varName, 0)
	assert.Equal(t, val, int8(-128))

	val = GetEnvInt8(varName, 2)
	assert.Equal(t, val, int8(-128))

	// Test with max value

	os.Setenv(varName, "127")

	val = GetEnvInt8(varName, 0)
	assert.Equal(t, val, int8(127))

	val = GetEnvInt8(varName, 2)
	assert.Equal(t, val, int8(127))

	// Test with overflow

	os.Setenv(varName, "128")

	val = GetEnvInt8(varName, 0)
	assert.Equal(t, val, int8(0))

	val = GetEnvInt8(varName, 2)
	assert.Equal(t, val, int8(2))
}

func TestGetEnvInt8WithWarning(t *testing.T) {
	var val int8
	var warning bool
	varName := "VAR_INT8_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvInt8WithWarning(varName, 0)
	assert.Equal(t, val, int8(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt8WithWarning(varName, 2)
	assert.Equal(t, val, int8(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvInt8WithWarning(varName, 0)
	assert.Equal(t, val, int8(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt8WithWarning(varName, 2)
	assert.Equal(t, val, int8(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvInt8WithWarning(varName, 0)
	assert.Equal(t, val, int8(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt8WithWarning(varName, 2)
	assert.Equal(t, val, int8(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvInt8WithWarning(varName, 0)
	assert.Equal(t, val, int8(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvInt8WithWarning(varName, 2)
	assert.Equal(t, val, int8(2))
	assert.Equal(t, warning, true)

	// Test with min value

	os.Setenv(varName, "-128")

	val, warning = GetEnvInt8WithWarning(varName, 0)
	assert.Equal(t, val, int8(-128))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt8WithWarning(varName, 2)
	assert.Equal(t, val, int8(-128))
	assert.Equal(t, warning, false)

	// Test with max value

	os.Setenv(varName, "127")

	val, warning = GetEnvInt8WithWarning(varName, 0)
	assert.Equal(t, val, int8(127))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt8WithWarning(varName, 2)
	assert.Equal(t, val, int8(127))
	assert.Equal(t, warning, false)

	// Test with overflow

	os.Setenv(varName, "128")

	val, warning = GetEnvInt8WithWarning(varName, 0)
	assert.Equal(t, val, int8(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvInt8WithWarning(varName, 2)
	assert.Equal(t, val, int8(2))
	assert.Equal(t, warning, true)
}

func TestGetEnvInt16(t *testing.T) {
	var val int16
	varName := "VAR_INT16_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvInt16(varName, 0)
	assert.Equal(t, val, int16(0))

	val = GetEnvInt16(varName, 2)
	assert.Equal(t, val, int16(2))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvInt16(varName, 0)
	assert.Equal(t, val, int16(0))

	val = GetEnvInt16(varName, 2)
	assert.Equal(t, val, int16(0))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvInt16(varName, 0)
	assert.Equal(t, val, int16(5))

	val = GetEnvInt16(varName, 2)
	assert.Equal(t, val, int16(5))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvInt16(varName, 0)
	assert.Equal(t, val, int16(0))

	val = GetEnvInt16(varName, 2)
	assert.Equal(t, val, int16(2))

	// Test with min value

	os.Setenv(varName, "-32768")

	val = GetEnvInt16(varName, 0)
	assert.Equal(t, val, int16(-32768))

	val = GetEnvInt16(varName, 2)
	assert.Equal(t, val, int16(-32768))

	// Test with max value

	os.Setenv(varName, "32767")

	val = GetEnvInt16(varName, 0)
	assert.Equal(t, val, int16(32767))

	val = GetEnvInt16(varName, 2)
	assert.Equal(t, val, int16(32767))

	// Test with overflow

	os.Setenv(varName, "32768")

	val = GetEnvInt16(varName, 0)
	assert.Equal(t, val, int16(0))

	val = GetEnvInt16(varName, 2)
	assert.Equal(t, val, int16(2))
}

func TestGetEnvInt16WithWarning(t *testing.T) {
	var val int16
	var warning bool
	varName := "VAR_INT16_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvInt16WithWarning(varName, 0)
	assert.Equal(t, val, int16(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt16WithWarning(varName, 2)
	assert.Equal(t, val, int16(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvInt16WithWarning(varName, 0)
	assert.Equal(t, val, int16(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt16WithWarning(varName, 2)
	assert.Equal(t, val, int16(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvInt16WithWarning(varName, 0)
	assert.Equal(t, val, int16(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt16WithWarning(varName, 2)
	assert.Equal(t, val, int16(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvInt16WithWarning(varName, 0)
	assert.Equal(t, val, int16(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvInt16WithWarning(varName, 2)
	assert.Equal(t, val, int16(2))
	assert.Equal(t, warning, true)

	// Test with min value

	os.Setenv(varName, "-32768")

	val, warning = GetEnvInt16WithWarning(varName, 0)
	assert.Equal(t, val, int16(-32768))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt16WithWarning(varName, 2)
	assert.Equal(t, val, int16(-32768))
	assert.Equal(t, warning, false)

	// Test with max value

	os.Setenv(varName, "32767")

	val, warning = GetEnvInt16WithWarning(varName, 0)
	assert.Equal(t, val, int16(32767))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt16WithWarning(varName, 2)
	assert.Equal(t, val, int16(32767))
	assert.Equal(t, warning, false)

	// Test with overflow

	os.Setenv(varName, "32768")

	val, warning = GetEnvInt16WithWarning(varName, 0)
	assert.Equal(t, val, int16(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvInt16WithWarning(varName, 2)
	assert.Equal(t, val, int16(2))
	assert.Equal(t, warning, true)
}

func TestGetEnvInt32(t *testing.T) {
	var val int32
	varName := "VAR_INT32_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvInt32(varName, 0)
	assert.Equal(t, val, int32(0))

	val = GetEnvInt32(varName, 2)
	assert.Equal(t, val, int32(2))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvInt32(varName, 0)
	assert.Equal(t, val, int32(0))

	val = GetEnvInt32(varName, 2)
	assert.Equal(t, val, int32(0))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvInt32(varName, 0)
	assert.Equal(t, val, int32(5))

	val = GetEnvInt32(varName, 2)
	assert.Equal(t, val, int32(5))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvInt32(varName, 0)
	assert.Equal(t, val, int32(0))

	val = GetEnvInt32(varName, 2)
	assert.Equal(t, val, int32(2))

	// Test with min value

	os.Setenv(varName, "-2147483648")

	val = GetEnvInt32(varName, 0)
	assert.Equal(t, val, int32(-2147483648))

	val = GetEnvInt32(varName, 2)
	assert.Equal(t, val, int32(-2147483648))

	// Test with max value

	os.Setenv(varName, "2147483647")

	val = GetEnvInt32(varName, 0)
	assert.Equal(t, val, int32(2147483647))

	val = GetEnvInt32(varName, 2)
	assert.Equal(t, val, int32(2147483647))

	// Test with overflow

	os.Setenv(varName, "2147483648")

	val = GetEnvInt32(varName, 0)
	assert.Equal(t, val, int32(0))

	val = GetEnvInt32(varName, 2)
	assert.Equal(t, val, int32(2))
}

func TestGetEnvInt32WithWarning(t *testing.T) {
	var val int32
	var warning bool
	varName := "VAR_INT32_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvInt32WithWarning(varName, 0)
	assert.Equal(t, val, int32(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt32WithWarning(varName, 2)
	assert.Equal(t, val, int32(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvInt32WithWarning(varName, 0)
	assert.Equal(t, val, int32(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt32WithWarning(varName, 2)
	assert.Equal(t, val, int32(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvInt32WithWarning(varName, 0)
	assert.Equal(t, val, int32(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt32WithWarning(varName, 2)
	assert.Equal(t, val, int32(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvInt32WithWarning(varName, 0)
	assert.Equal(t, val, int32(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvInt32WithWarning(varName, 2)
	assert.Equal(t, val, int32(2))
	assert.Equal(t, warning, true)

	// Test with min value

	os.Setenv(varName, "-2147483648")

	val, warning = GetEnvInt32WithWarning(varName, 0)
	assert.Equal(t, val, int32(-2147483648))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt32WithWarning(varName, 2)
	assert.Equal(t, val, int32(-2147483648))
	assert.Equal(t, warning, false)

	// Test with max value

	os.Setenv(varName, "2147483647")

	val, warning = GetEnvInt32WithWarning(varName, 0)
	assert.Equal(t, val, int32(2147483647))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt32WithWarning(varName, 2)
	assert.Equal(t, val, int32(2147483647))
	assert.Equal(t, warning, false)

	// Test with overflow

	os.Setenv(varName, "2147483648")

	val, warning = GetEnvInt32WithWarning(varName, 0)
	assert.Equal(t, val, int32(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvInt32WithWarning(varName, 2)
	assert.Equal(t, val, int32(2))
	assert.Equal(t, warning, true)
}

func TestGetEnvInt64(t *testing.T) {
	var val int64
	varName := "VAR_INT64_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvInt64(varName, 0)
	assert.Equal(t, val, int64(0))

	val = GetEnvInt64(varName, 2)
	assert.Equal(t, val, int64(2))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvInt64(varName, 0)
	assert.Equal(t, val, int64(0))

	val = GetEnvInt64(varName, 2)
	assert.Equal(t, val, int64(0))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvInt64(varName, 0)
	assert.Equal(t, val, int64(5))

	val = GetEnvInt64(varName, 2)
	assert.Equal(t, val, int64(5))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvInt64(varName, 0)
	assert.Equal(t, val, int64(0))

	val = GetEnvInt64(varName, 2)
	assert.Equal(t, val, int64(2))

	// Test with min value

	os.Setenv(varName, "-9223372036854775808")

	val = GetEnvInt64(varName, 0)
	assert.Equal(t, val, int64(-9223372036854775808))

	val = GetEnvInt64(varName, 2)
	assert.Equal(t, val, int64(-9223372036854775808))

	// Test with max value

	os.Setenv(varName, "9223372036854775807")

	val = GetEnvInt64(varName, 0)
	assert.Equal(t, val, int64(9223372036854775807))

	val = GetEnvInt64(varName, 2)
	assert.Equal(t, val, int64(9223372036854775807))

	// Test with overflow

	os.Setenv(varName, "9223372036854775808")

	val = GetEnvInt64(varName, 0)
	assert.Equal(t, val, int64(0))

	val = GetEnvInt64(varName, 2)
	assert.Equal(t, val, int64(2))
}

func TestGetEnvInt64WithWarning(t *testing.T) {
	var val int64
	var warning bool
	varName := "VAR_INT64_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvInt64WithWarning(varName, 0)
	assert.Equal(t, val, int64(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt64WithWarning(varName, 2)
	assert.Equal(t, val, int64(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvInt64WithWarning(varName, 0)
	assert.Equal(t, val, int64(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt64WithWarning(varName, 2)
	assert.Equal(t, val, int64(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvInt64WithWarning(varName, 0)
	assert.Equal(t, val, int64(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt64WithWarning(varName, 2)
	assert.Equal(t, val, int64(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvInt64WithWarning(varName, 0)
	assert.Equal(t, val, int64(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvInt64WithWarning(varName, 2)
	assert.Equal(t, val, int64(2))
	assert.Equal(t, warning, true)

	// Test with min value

	os.Setenv(varName, "-9223372036854775808")

	val, warning = GetEnvInt64WithWarning(varName, 0)
	assert.Equal(t, val, int64(-9223372036854775808))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt64WithWarning(varName, 2)
	assert.Equal(t, val, int64(-9223372036854775808))
	assert.Equal(t, warning, false)

	// Test with max value

	os.Setenv(varName, "9223372036854775807")

	val, warning = GetEnvInt64WithWarning(varName, 0)
	assert.Equal(t, val, int64(9223372036854775807))
	assert.Equal(t, warning, false)

	val, warning = GetEnvInt64WithWarning(varName, 2)
	assert.Equal(t, val, int64(9223372036854775807))
	assert.Equal(t, warning, false)

	// Test with overflow

	os.Setenv(varName, "9223372036854775808")

	val, warning = GetEnvInt64WithWarning(varName, 0)
	assert.Equal(t, val, int64(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvInt64WithWarning(varName, 2)
	assert.Equal(t, val, int64(2))
	assert.Equal(t, warning, true)
}
