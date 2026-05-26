# Unix WC tool using Go

A WC tool named `ccwc` built from scratch using Go.  
Analyzes input streams and provides counts for:

- Bytes
- Lines
- Words
- Runes

## Documentation

- [**Overview**](#overview)
- [**Project Structure**](#project-structure)
- [**Example**](#example)
- [**Setup**](#setup)
- [**Build Process**](#build-process)
- [**License**](#license)


## Overview

This project is a custom implementation of the Unix `wc` utility built from scratch. The goal is to understand how file analysis and stream processing work at a low level.

The tool works in two stages:

#### Reader

Detects the input mode and provides a stream to the Analyzer.

- `Terminal mode`: reads from a file path provided as a cli argument
- `Stdin mode`: reads from piped standard input

#### Analyzer

Reads the input stream in 32kB chunks and computes results for each selected option.

- `-c`: counts total bytes in the stream
- `-l`: counts newline characters
- `-w`: counts whitespace-delimited tokens
- `-m`: counts Unicode code points using UTF-8 decoding


## Project structure

```
folder: challenge-01/
├── cmd/
|   └── ccwc.go          # entry point
│
├── config/
│   └── config.go        # defines central configuration
│
├── analyzer/
│   ├── reader.go        # handles reading of input 
│   ├── analyzer.go      # analyzes required options
│   └── errors.go        # predefined errors
|
├── test.txt             # test file 
|
├── diagrams/            # [WIP]
|
├── file_test.go         # [WIP]

```

## Example

Input:

```bash
ccwc -w -l test.txt
```

Output:

```bash
map[Lines:7145 Words:58164]
```

## Setup

### Prerequisites

- [Go](https://go.dev/doc/install) 1.23 or later

- follow the steps [here](https://github.com/priyanshsao/coding-challenges-go/blob/main/README.md#setup) to setup parent repository.

### Change Directory

```bash
cd challenge-01/
```

### Run program

```bash
go run ./cmd/ccwc.go
```

### Build binary

```bash
go build -o ./bin/ccwc ./cmd/ccwc.go
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

- Main steps for building this project are tracked in [epic issue](https://github.com/priyanshsao/coding-challenges-go/issues/10).

## License

Licensed under [MIT](https://github.com/priyanshsao/coding-challenges-go/blob/main/LICENSE).