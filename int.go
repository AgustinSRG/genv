// Environment variables as int

package genv

import (
	"os"
	"strconv"
)

// Gets environment variable as integer.
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvInt(variableName string, defaultValue int) int {
	v, _ := GetEnvIntWithWarning(variableName, defaultValue)
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
func GetEnvIntWithWarning(variableName string, defaultValue int) (val int, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.Atoi(vString)

	if e != nil {
		return defaultValue, true
	}

	return v, false
}

// Gets environment variable as integer (int8).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvInt8(variableName string, defaultValue int8) int8 {
	v, _ := GetEnvInt8WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as integer (int8).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvInt8WithWarning(variableName string, defaultValue int8) (val int8, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseInt(vString, 0, 8)

	if e != nil {
		return defaultValue, true
	}

	return int8(v), false
}

// Gets environment variable as integer (int16).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvInt16(variableName string, defaultValue int16) int16 {
	v, _ := GetEnvInt16WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as integer (int16).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvInt16WithWarning(variableName string, defaultValue int16) (val int16, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseInt(vString, 0, 16)

	if e != nil {
		return defaultValue, true
	}

	return int16(v), false
}

// Gets environment variable as integer (int32).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvInt32(variableName string, defaultValue int32) int32 {
	v, _ := GetEnvInt32WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as integer (int32).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvInt32WithWarning(variableName string, defaultValue int32) (val int32, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseInt(vString, 0, 32)

	if e != nil {
		return defaultValue, true
	}

	return int32(v), false
}

// Gets environment variable as integer (int64).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvInt64(variableName string, defaultValue int64) int64 {
	v, _ := GetEnvInt64WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as integer (int64).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvInt64WithWarning(variableName string, defaultValue int64) (val int64, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseInt(vString, 0, 64)

	if e != nil {
		return defaultValue, true
	}

	return v, false
}
