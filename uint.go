// Environment variables as uint

package genv

import (
	"os"
	"strconv"
)

// Gets environment variable as unsigned integer.
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvUint(variableName string, defaultValue uint) uint {
	v, _ := GetEnvUintWithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as integer.
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvUintWithWarning(variableName string, defaultValue uint) (val uint, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseUint(vString, 0, 0)

	if e != nil {
		return defaultValue, true
	}

	return uint(v), false
}

// Gets environment variable as unsigned integer (uint8).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvUint8(variableName string, defaultValue uint8) uint8 {
	v, _ := GetEnvUint8WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as unsigned integer (uint8).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvUint8WithWarning(variableName string, defaultValue uint8) (val uint8, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseUint(vString, 0, 8)

	if e != nil {
		return defaultValue, true
	}

	return uint8(v), false
}

// Gets environment variable as unsigned integer (uint16).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvUint16(variableName string, defaultValue uint16) uint16 {
	v, _ := GetEnvUint16WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as unsigned integer (uint16).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvUint16WithWarning(variableName string, defaultValue uint16) (val uint16, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseUint(vString, 0, 16)

	if e != nil {
		return defaultValue, true
	}

	return uint16(v), false
}

// Gets environment variable as unsigned integer (uint32).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvUint32(variableName string, defaultValue uint32) uint32 {
	v, _ := GetEnvUint32WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as unsigned integer (uint32).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvUint32WithWarning(variableName string, defaultValue uint32) (val uint32, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseUint(vString, 0, 32)

	if e != nil {
		return defaultValue, true
	}

	return uint32(v), false
}

// Gets environment variable as unsigned integer (uint64).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvUint64(variableName string, defaultValue uint64) uint64 {
	v, _ := GetEnvUint64WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as unsigned integer (uint64).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvUint64WithWarning(variableName string, defaultValue uint64) (val uint64, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseUint(vString, 0, 64)

	if e != nil {
		return defaultValue, true
	}

	return v, false
}
