# Utility library to read and parse environment variables (Go)

This is a very simple library to read environment variables and parse them to multiple types.

[Documentation](https://pkg.go.dev/github.com/AgustinSRG/genv)

## Installation

To install the library in your project, run:

```sh
go get github.com/AgustinSRG/genv
```

## Usage

This library provides several functions that get environment variables and automatically parse them to different types. You can also set default values in case the variable is empty or missing.

Here is en example usage

```go
package main

import (
    "fmt"
    // Import the module
    "github.com/AgustinSRG/genv"
)

func main() {
    // You can get string variables
    // If not set, you get the default value, passed as the second argument
    boolVar := genv.GetEnvBool("TEST_STR_VAR", "default value")
    fmt.Printf("TEST_STR_VAR = %v\n", boolVar)

    // You can parse variables into boolean
    // If they are set to TRUE or YES, you get true.
    // If they are set to FALSE or NO, you get false.
    // If they are not set to any of the above, you get the default value, passed as the second argument
    strVar := genv.GetEnvBool("TEST_BOOL_VAR", false)
    fmt.Printf("Parsed TEST_BOOL_VAR = %v\n", strVar)

    // You can parse variables into integer
    // If not set, or invalid, you get the default value, passed as the second argument
    port := genv.GetEnvInt("PORT", 80)
    fmt.Printf("PORT = %v\n", port)

    // For types like integers, where the value can be invalid,
    // you can use the alternative functions ending with "WithWarning"
    // They do the same, but also return a boolean flag, set to true
    // only if the value is set, but invalid
    port, warning := genv.GetEnvIntWithWarning("PORT", 80)
    if warning {
        fmt.Println("[Warning] PORT has an invalid integer value, using the default value.")
    }
    fmt.Printf("PORT = %v\n", port)

    // Same for the rest of the types. Check the documentation for more
}
```

## Build the library

To install dependencies, run:

```sh
go get .
```

To build the code, run:

```sh
go build .
```

## Run linter

To run the code linter, run:

```sh
golangci-lint run
```

## Run tests

In order to run the tests for this library, run:

```sh
go test -v
```
