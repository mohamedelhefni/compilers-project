# Simple Parser in Go

This repository contains a simple lexer and parser written in Go. The parser can handle basic assignments and if statements.

## Lexer (Lexical Analysis)

The lexer (`lexer.go`) tokenizes input strings into a sequence of tokens. The supported token types include:

- `INTEGER`: Integer literals
- `PLUS`, `MINUS`, `MULTIPLY`, `DIVIDE`: Arithmetic operators
- `LPAREN`, `RPAREN`: Left and right parentheses
- `ASSIGN`: Assignment operator
- `ID`: Identifier (variable name)
- `IF`, `ELSE`: If statement and else keyword
- `GREATER_THAN`: Greater than operator
- `LBRACE`, `RBRACE`: Left and right curly braces

## Parser (Syntax Analysis)

The parser (`parser.go`) processes a list of tokens produced by the lexer. It can handle simple assignments and if statements. The grammar includes:

- Statements: Assignments and if statements
- Expressions: Arithmetic expressions and comparisons
- Terms and factors: Components of expressions


