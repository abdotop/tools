# Key Package

The `key` package provides utilities to generate random passwords with customizable complexity.

## Features

- Generate random passwords with a specified length.
- Optionally include numbers and special characters for increased password complexity.


## Installation

To install the package, you can use `go get`:
```bash
go get github.com/abdotop/tools/key
```

## Usage

To use the `key` package, you need to import it into your Go project:
```go
import "github.com/abdotop/tools/key"
```

### Generating a Password

You can generate a password by calling the `GenerateKey` function. Here is an example of how to generate a 12-character password that includes numbers and special characters:

```go
password := key.Generate(12, true, true)
fmt.Println("Generated Password:", password)
```

## License

This package is licensed under the MIT License. See the LICENSE file for more details.