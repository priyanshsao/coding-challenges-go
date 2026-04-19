# JSON Parser using Go

## Documentation
- [**Overview**](#overview)
- [**Project Structure**](#project-structure)
- [**Setup**](#setup)
- [**Build Process**](#build-process)
- [**License**](#license)


## Overview
This project is a custom implementation of a JSON parser built from scratch. The goal is to understand how structured data (like JSON) is processed internally.

The parser works in multiple stages:

- #### Lexer (tokenizer) 
    reads the json input character by character and converts it into a sequence of tokens. Tokens can be of following type
    - structural tokens:`{`,`}`,`[`,`]` 
    - strings: `"hello"`
    - numbers: `ex-12.32`
    - literals: `true`/`false`/`null`

- #### Parser
    uses tokens provided by lexer and analyzes their structure. 
    
    It ensures the JSON follows correct syntax rules (like matching brackets and proper key-value pairs).
    
    Based on this, it builds corresponding data structures such as arrays, objects (maps), strings, and numbers.


## Project structure

```
folder: challenge-02/
├── tokens.go       # define tokens   
├── lexer.go        # converts input into tokens
├── errors.go       # introduce predefined errors
├── parser.go       # parses tokens according to JSON rules
├── main.go         # tests against files in 'test' folder

├── diagrams/       # flow diagrams for each layer of project
├── tests/          # test files to test success of each step

├── lexer_test.go   # unit tests for lexer
├── parser_test.go  # unit tests for parser

```