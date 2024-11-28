// Tests

package genv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvUint(t *testing.T) {
	var val uint
	varName := "VAR_UINT_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvUint(varName, 0)
	assert.Equal(t, val, uint(0))

	val = GetEnvUint(varName, 2)
	assert.Equal(t, val, uint(2))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvUint(varName, 0)
	assert.Equal(t, val, uint(5))

	val = GetEnvUint(varName, 2)
	assert.Equal(t, val, uint(5))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvUint(varName, 0)
	assert.Equal(t, val, uint(0))

	val = GetEnvUint(varName, 2)
	assert.Equal(t, val, uint(0))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvUint(varName, 0)
	assert.Equal(t, val, uint(0))

	val = GetEnvUint(varName, 2)
	assert.Equal(t, val, uint(2))
}

func TestGetEnvUintWithWarning(t *testing.T) {
	var val uint
	var warning bool
	varName := "VAR_UINT_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvUintWithWarning(varName, 0)
	assert.Equal(t, val, uint(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUintWithWarning(varName, 2)
	assert.Equal(t, val, uint(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvUintWithWarning(varName, 0)
	assert.Equal(t, val, uint(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUintWithWarning(varName, 2)
	assert.Equal(t, val, uint(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvUintWithWarning(varName, 0)
	assert.Equal(t, val, uint(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUintWithWarning(varName, 2)
	assert.Equal(t, val, uint(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvUintWithWarning(varName, 0)
	assert.Equal(t, val, uint(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvUintWithWarning(varName, 2)
	assert.Equal(t, val, uint(2))
	assert.Equal(t, warning, true)
}

func TestGetEnvUint8(t *testing.T) {
	var val uint8
	varName := "VAR_UINT8_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvUint8(varName, 0)
	assert.Equal(t, val, uint8(0))

	val = GetEnvUint8(varName, 2)
	assert.Equal(t, val, uint8(2))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvUint8(varName, 0)
	assert.Equal(t, val, uint8(0))

	val = GetEnvUint8(varName, 2)
	assert.Equal(t, val, uint8(0))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvUint8(varName, 0)
	assert.Equal(t, val, uint8(5))

	val = GetEnvUint8(varName, 2)
	assert.Equal(t, val, uint8(5))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvUint8(varName, 0)
	assert.Equal(t, val, uint8(0))

	val = GetEnvUint8(varName, 2)
	assert.Equal(t, val, uint8(2))

	// Test with max value

	os.Setenv(varName, "255")

	val = GetEnvUint8(varName, 0)
	assert.Equal(t, val, uint8(255))

	val = GetEnvUint8(varName, 2)
	assert.Equal(t, val, uint8(255))

	// Test with overflow

	os.Setenv(varName, "256")

	val = GetEnvUint8(varName, 0)
	assert.Equal(t, val, uint8(0))

	val = GetEnvUint8(varName, 2)
	assert.Equal(t, val, uint8(2))
}

func TestGetEnvUint8WithWarning(t *testing.T) {
	var val uint8
	var warning bool
	varName := "VAR_UINT8_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvUint8WithWarning(varName, 0)
	assert.Equal(t, val, uint8(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint8WithWarning(varName, 2)
	assert.Equal(t, val, uint8(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvUint8WithWarning(varName, 0)
	assert.Equal(t, val, uint8(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint8WithWarning(varName, 2)
	assert.Equal(t, val, uint8(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvUint8WithWarning(varName, 0)
	assert.Equal(t, val, uint8(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint8WithWarning(varName, 2)
	assert.Equal(t, val, uint8(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvUint8WithWarning(varName, 0)
	assert.Equal(t, val, uint8(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvUint8WithWarning(varName, 2)
	assert.Equal(t, val, uint8(2))
	assert.Equal(t, warning, true)

	// Test with max value

	os.Setenv(varName, "255")

	val, warning = GetEnvUint8WithWarning(varName, 0)
	assert.Equal(t, val, uint8(255))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint8WithWarning(varName, 2)
	assert.Equal(t, val, uint8(255))
	assert.Equal(t, warning, false)

	// Test with overflow

	os.Setenv(varName, "256")

	val, warning = GetEnvUint8WithWarning(varName, 0)
	assert.Equal(t, val, uint8(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvUint8WithWarning(varName, 2)
	assert.Equal(t, val, uint8(2))
	assert.Equal(t, warning, true)
}

func TestGetEnvUint16(t *testing.T) {
	var val uint16
	varName := "VAR_UINT16_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvUint16(varName, 0)
	assert.Equal(t, val, uint16(0))

	val = GetEnvUint16(varName, 2)
	assert.Equal(t, val, uint16(2))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvUint16(varName, 0)
	assert.Equal(t, val, uint16(0))

	val = GetEnvUint16(varName, 2)
	assert.Equal(t, val, uint16(0))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvUint16(varName, 0)
	assert.Equal(t, val, uint16(5))

	val = GetEnvUint16(varName, 2)
	assert.Equal(t, val, uint16(5))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvUint16(varName, 0)
	assert.Equal(t, val, uint16(0))

	val = GetEnvUint16(varName, 2)
	assert.Equal(t, val, uint16(2))

	// Test with max value

	os.Setenv(varName, "65535")

	val = GetEnvUint16(varName, 0)
	assert.Equal(t, val, uint16(65535))

	val = GetEnvUint16(varName, 2)
	assert.Equal(t, val, uint16(65535))

	// Test with overflow

	os.Setenv(varName, "65536")

	val = GetEnvUint16(varName, 0)
	assert.Equal(t, val, uint16(0))

	val = GetEnvUint16(varName, 2)
	assert.Equal(t, val, uint16(2))
}

func TestGetEnvUint16WithWarning(t *testing.T) {
	var val uint16
	var warning bool
	varName := "VAR_UINT16_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvUint16WithWarning(varName, 0)
	assert.Equal(t, val, uint16(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint16WithWarning(varName, 2)
	assert.Equal(t, val, uint16(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvUint16WithWarning(varName, 0)
	assert.Equal(t, val, uint16(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint16WithWarning(varName, 2)
	assert.Equal(t, val, uint16(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvUint16WithWarning(varName, 0)
	assert.Equal(t, val, uint16(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint16WithWarning(varName, 2)
	assert.Equal(t, val, uint16(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvUint16WithWarning(varName, 0)
	assert.Equal(t, val, uint16(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvUint16WithWarning(varName, 2)
	assert.Equal(t, val, uint16(2))
	assert.Equal(t, warning, true)

	// Test with max value

	os.Setenv(varName, "65535")

	val, warning = GetEnvUint16WithWarning(varName, 0)
	assert.Equal(t, val, uint16(65535))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint16WithWarning(varName, 2)
	assert.Equal(t, val, uint16(65535))
	assert.Equal(t, warning, false)

	// Test with overflow

	os.Setenv(varName, "65536")

	val, warning = GetEnvUint16WithWarning(varName, 0)
	assert.Equal(t, val, uint16(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvUint16WithWarning(varName, 2)
	assert.Equal(t, val, uint16(2))
	assert.Equal(t, warning, true)
}

func TestGetEnvUint32(t *testing.T) {
	var val uint32
	varName := "VAR_UINT32_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvUint32(varName, 0)
	assert.Equal(t, val, uint32(0))

	val = GetEnvUint32(varName, 2)
	assert.Equal(t, val, uint32(2))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvUint32(varName, 0)
	assert.Equal(t, val, uint32(0))

	val = GetEnvUint32(varName, 2)
	assert.Equal(t, val, uint32(0))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvUint32(varName, 0)
	assert.Equal(t, val, uint32(5))

	val = GetEnvUint32(varName, 2)
	assert.Equal(t, val, uint32(5))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvUint32(varName, 0)
	assert.Equal(t, val, uint32(0))

	val = GetEnvUint32(varName, 2)
	assert.Equal(t, val, uint32(2))

	// Test with max value

	os.Setenv(varName, "4294967295")

	val = GetEnvUint32(varName, 0)
	assert.Equal(t, val, uint32(4294967295))

	val = GetEnvUint32(varName, 2)
	assert.Equal(t, val, uint32(4294967295))

	// Test with overflow

	os.Setenv(varName, "4294967296")

	val = GetEnvUint32(varName, 0)
	assert.Equal(t, val, uint32(0))

	val = GetEnvUint32(varName, 2)
	assert.Equal(t, val, uint32(2))
}

func TestGetEnvUint32WithWarning(t *testing.T) {
	var val uint32
	var warning bool
	varName := "VAR_UINT32_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvUint32WithWarning(varName, 0)
	assert.Equal(t, val, uint32(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint32WithWarning(varName, 2)
	assert.Equal(t, val, uint32(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvUint32WithWarning(varName, 0)
	assert.Equal(t, val, uint32(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint32WithWarning(varName, 2)
	assert.Equal(t, val, uint32(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvUint32WithWarning(varName, 0)
	assert.Equal(t, val, uint32(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint32WithWarning(varName, 2)
	assert.Equal(t, val, uint32(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvUint32WithWarning(varName, 0)
	assert.Equal(t, val, uint32(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvUint32WithWarning(varName, 2)
	assert.Equal(t, val, uint32(2))
	assert.Equal(t, warning, true)

	// Test with max value

	os.Setenv(varName, "4294967295")

	val, warning = GetEnvUint32WithWarning(varName, 0)
	assert.Equal(t, val, uint32(4294967295))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint32WithWarning(varName, 2)
	assert.Equal(t, val, uint32(4294967295))
	assert.Equal(t, warning, false)

	// Test with overflow

	os.Setenv(varName, "4294967296")

	val, warning = GetEnvUint32WithWarning(varName, 0)
	assert.Equal(t, val, uint32(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvUint32WithWarning(varName, 2)
	assert.Equal(t, val, uint32(2))
	assert.Equal(t, warning, true)
}

func TestGetEnvUint64(t *testing.T) {
	var val uint64
	varName := "VAR_UINT64_1"

	// Test empty

	os.Setenv(varName, "")

	val = GetEnvUint64(varName, 0)
	assert.Equal(t, val, uint64(0))

	val = GetEnvUint64(varName, 2)
	assert.Equal(t, val, uint64(2))

	// Test zero

	os.Setenv(varName, "0")

	val = GetEnvUint64(varName, 0)
	assert.Equal(t, val, uint64(0))

	val = GetEnvUint64(varName, 2)
	assert.Equal(t, val, uint64(0))

	// Test valid value

	os.Setenv(varName, "5")

	val = GetEnvUint64(varName, 0)
	assert.Equal(t, val, uint64(5))

	val = GetEnvUint64(varName, 2)
	assert.Equal(t, val, uint64(5))

	// Test invalid value

	os.Setenv(varName, "ab")

	val = GetEnvUint64(varName, 0)
	assert.Equal(t, val, uint64(0))

	val = GetEnvUint64(varName, 2)
	assert.Equal(t, val, uint64(2))

	// Test with max value

	os.Setenv(varName, "18446744073709551615")

	val = GetEnvUint64(varName, 0)
	assert.Equal(t, val, uint64(18446744073709551615))

	val = GetEnvUint64(varName, 2)
	assert.Equal(t, val, uint64(18446744073709551615))

	// Test with overflow

	os.Setenv(varName, "18446744073709551616")

	val = GetEnvUint64(varName, 0)
	assert.Equal(t, val, uint64(0))

	val = GetEnvUint64(varName, 2)
	assert.Equal(t, val, uint64(2))
}

func TestGetEnvUint64WithWarning(t *testing.T) {
	var val uint64
	var warning bool
	varName := "VAR_UINT64_2"

	// Test empty

	os.Setenv(varName, "")

	val, warning = GetEnvUint64WithWarning(varName, 0)
	assert.Equal(t, val, uint64(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint64WithWarning(varName, 2)
	assert.Equal(t, val, uint64(2))
	assert.Equal(t, warning, false)

	// Test zero

	os.Setenv(varName, "0")

	val, warning = GetEnvUint64WithWarning(varName, 0)
	assert.Equal(t, val, uint64(0))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint64WithWarning(varName, 2)
	assert.Equal(t, val, uint64(0))
	assert.Equal(t, warning, false)

	// Test valid value

	os.Setenv(varName, "5")

	val, warning = GetEnvUint64WithWarning(varName, 0)
	assert.Equal(t, val, uint64(5))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint64WithWarning(varName, 2)
	assert.Equal(t, val, uint64(5))
	assert.Equal(t, warning, false)

	// Test invalid value

	os.Setenv(varName, "ab")

	val, warning = GetEnvUint64WithWarning(varName, 0)
	assert.Equal(t, val, uint64(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvUint64WithWarning(varName, 2)
	assert.Equal(t, val, uint64(2))
	assert.Equal(t, warning, true)

	// Test with max value

	os.Setenv(varName, "18446744073709551615")

	val, warning = GetEnvUint64WithWarning(varName, 0)
	assert.Equal(t, val, uint64(18446744073709551615))
	assert.Equal(t, warning, false)

	val, warning = GetEnvUint64WithWarning(varName, 2)
	assert.Equal(t, val, uint64(18446744073709551615))
	assert.Equal(t, warning, false)

	// Test with overflow

	os.Setenv(varName, "18446744073709551616")

	val, warning = GetEnvUint64WithWarning(varName, 0)
	assert.Equal(t, val, uint64(0))
	assert.Equal(t, warning, true)

	val, warning = GetEnvUint64WithWarning(varName, 2)
	assert.Equal(t, val, uint64(2))
	assert.Equal(t, warning, true)
}
