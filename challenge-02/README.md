# JSON Parser using Go

A JSON parser built from scratch in Go. Parses and validates JSON strings following the [JSON spec](https://www.json.org/json-en.html).

## Documentation

- [**Overview**](#overview)
- [**Project Structure**](#project-structure)
- [**Example**](#example)
- [**Setup**](#setup)
- [**Build Process**](#build-process)
- [**License**](#license)


## Overview

This project is a custom implementation of a JSON parser built from scratch. The goal is to understand how structured data (like JSON) is processed internally.

The parser works in multiple stages:

#### Lexer (tokenizer) 

Reads the json input character by character and converts it into a sequence of tokens. Tokens can be of following type

- structural tokens:`{`,`}`,`[`,`]` 
- strings: `"hello"`
- numbers: `ex-12.32`
- literals: `true`/`false`/`null`

#### Parser

- Uses tokens provided by lexer and analyzes their structure. 

- It ensures the JSON follows correct syntax rules (like matching brackets and proper key-value pairs).

- Based on this, it builds corresponding data structures such as arrays, objects (maps), strings, and numbers.


## Project structure

```
folder: challenge-02/
├── cmd/
│   └── main.go                 # Runs all the tests
│
├── internal/
│   ├── define/
│   │   └── tokens.go           # Token definitions
│   │
│   ├── lexer/
│   │   ├── define.go           # Lexer definition
│   │   ├── lexer.go            # Lexer methods
│   │   ├── lexer_test.go       # Unit tests for lexer
│   │   └── utils.go            # Helper functions
│   │
│   └── parser/
│       ├── define.go           # Parser definition
│       ├── parser.go           # Parser methods
│       ├── parser_test.go      # Unit tests for parser
│       └── errors.go           # Predefined parser errors
│
├── json/                       # main package to import
│   ├── factory.go              # Factory functions for internals
│   └── interface.go            # Parser interface definitions
│
├── diagrams/                   # Flow diagrams for each project layer
│
└── tests/                      # Test files for validating each step
```

## Example

Input:

```json
{"name": "Priyansh", "age": 37}
```

Output:

```go
map[string]any{
"name": "Priyansh",
"age": 37,
}
```

## Setup

### Prerequisites

- [Go](https://go.dev/doc/install) 1.23 or later

- follow the steps [here](https://github.com/priyanshsao/coding-challenges-go/blob/main/README.md#setup) to setup parent repository.

### Change Directory

```bash
cd challenge-02/
```

### Run program

```bash
go run ./cmd/main.go
```

### Run unit tests

```bash
go test -v
```

### Get test coverage profile

```bash
go test ./ -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
start coverage.html
```

> [!TIP]
> Clean up coverage files using this
> `rm coverage.out coverage.html`

## Build Process
- Each and every feature and bugs are tracked using github issues.

- Main steps for building this project are tracked in [epic issue](https://github.com/priyanshsao/coding-challenges-go/issues/11).

## License

Licensed under [MIT](https://github.com/priyanshsao/coding-challenges-go/blob/main/LICENSE).