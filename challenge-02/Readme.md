# JSON Parser using Go

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
├── tokens.go       # define tokens   
├── lexer.go        # converts input into tokens
├── errors.go       # introduce predefined errors
├── parser.go       # parses tokens according to JSON rules
├── main.go         # tests against files in 'test' folder
|
├── diagrams/       # flow diagrams for each layer of project
├── tests/          # test files to test success of each step
|
├── lexer_test.go   # unit tests for lexer
├── parser_test.go  # unit tests for parser

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

#### Prerequisites

- [Go](https://go.dev/doc/install) 1.23 or later

- follow the steps [here](https://github.com/priyanshsao/coding-challenges-go/blob/main/README.md#setup) to setup parent repository.

### Change Directory

```bash
cd challenge-02/
```

### Run program

```bash
go run ./
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

Licensed under [MIT](LICENSE).