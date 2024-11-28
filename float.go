// Environment variables as int

package genv

import (
	"os"
	"strconv"
)

// Gets environment variable as float (float32).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvFloat32(variableName string, defaultValue float32) float32 {
	v, _ := GetEnvFloat32WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as float (float32).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvFloat32WithWarning(variableName string, defaultValue float32) (val float32, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseFloat(vString, 32)

	if e != nil {
		return defaultValue, true
	}

	return float32(v), false
}

// Gets environment variable as float (float64).
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the parsed value of the variable.
// If empty, or the value cannot be parsed, the default value will be returned
func GetEnvFloat64(variableName string, defaultValue float64) float64 {
	v, _ := GetEnvFloat64WithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as float (float64).
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvFloat64WithWarning(variableName string, defaultValue float64) (val float64, warning bool) {
	vString := os.Getenv(variableName)

	if vString == "" {
		return defaultValue, false
	}

	v, e := strconv.ParseFloat(vString, 64)

	if e != nil {
		return defaultValue, true
	}

	return v, false
}
