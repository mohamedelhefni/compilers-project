package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType string

const (
	INTEGER      TokenType = "INTEGER"
	PLUS         TokenType = "PLUS"
	MINUS        TokenType = "MINUS"
	MULTIPLY     TokenType = "MULTIPLY"
	DIVIDE       TokenType = "DIVIDE"
	LPAREN       TokenType = "LPAREN"
	RPAREN       TokenType = "RPAREN"
	ASSIGN       TokenType = "ASSIGN"
	ID           TokenType = "ID"
	IF           TokenType = "IF"
	ELSE         TokenType = "ELSE"
	ENDIF        TokenType = "ENDIF"
	THEN         TokenType = "THEN"
	GREATER_THAN TokenType = "GREATER_THAN"
	LBRACE       TokenType = "LBRACE"
	RBRACE       TokenType = "RBRACE"
	LBRACKET     TokenType = "LBRACKET"
	RBRACKET     TokenType = "RBRACKET"
	COMMA        TokenType = "COMMA"
	DOT          TokenType = "DOT"
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("(%v, %v)", t.Type, t.Value)
}

type Lexer struct {
	input    string
	position int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, position: 0}
}

func (l *Lexer) tokenize() []Token {
	var tokens []Token

	for l.position < len(l.input) {
		currentChar := l.input[l.position]

		if unicode.IsDigit(rune(currentChar)) {
			tokens = append(tokens, Token{Type: INTEGER, Value: l.readNumber()})
		} else if unicode.IsSpace(rune(currentChar)) {
			l.position++
		} else {
			switch currentChar {
			case '+':
				tokens = append(tokens, Token{Type: PLUS, Value: string(currentChar)})
				l.position++
			case '-':
				tokens = append(tokens, Token{Type: MINUS, Value: string(currentChar)})
				l.position++
			case '*':
				tokens = append(tokens, Token{Type: MULTIPLY, Value: string(currentChar)})
				l.position++
			case '/':
				tokens = append(tokens, Token{Type: DIVIDE, Value: string(currentChar)})
				l.position++
			case '(':
				tokens = append(tokens, Token{Type: LPAREN, Value: string(currentChar)})
				l.position++
			case ')':
				tokens = append(tokens, Token{Type: RPAREN, Value: string(currentChar)})
				l.position++
			case '=':
				tokens = append(tokens, Token{Type: ASSIGN, Value: string(currentChar)})
				l.position++
			case '>':
				tokens = append(tokens, Token{Type: GREATER_THAN, Value: string(currentChar)})
				l.position++
			case '{':
				tokens = append(tokens, Token{Type: LBRACE, Value: string(currentChar)})
				l.position++
			case '}':
				tokens = append(tokens, Token{Type: RBRACE, Value: string(currentChar)})
				l.position++
			default:
				if unicode.IsLetter(rune(currentChar)) {
					identifier := l.readIdentifier()
					switch identifier {
					case "if":
						tokens = append(tokens, Token{Type: IF, Value: identifier})
					case "else":
						tokens = append(tokens, Token{Type: ELSE, Value: identifier})
					default:
						tokens = append(tokens, Token{Type: ID, Value: identifier})
					}
				} else {
					fmt.Printf("Invalid character: %c\n", currentChar)
					return nil
				}
			}
		}
	}

	return tokens
}

func (l *Lexer) readNumber() string {
	var result strings.Builder
	for l.position < len(l.input) && unicode.IsDigit(rune(l.input[l.position])) {
		result.WriteRune(rune(l.input[l.position]))
		l.position++
	}
	return result.String()
}

func (l *Lexer) readIdentifier() string {
	var result strings.Builder
	for l.position < len(l.input) && (unicode.IsLetter(rune(l.input[l.position])) || unicode.IsDigit(rune(l.input[l.position]))) {
		result.WriteRune(rune(l.input[l.position]))
		l.position++
	}
	return result.String()
}

type Parser struct {
	tokens  []Token
	current int
}

type ParseError struct {
	Message string
}

func (e *ParseError) Error() string {
	return e.Message
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) parse() error {
	for p.current < len(p.tokens) {
		if err := p.statement(); err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) statement() error {
	if p.match(ID) {
		return p.assignment()
	} else if p.match(IF) {
		p.match(LPAREN)
		p.expression()
		p.match(RPAREN)
		p.match(LBRACE)
		if err := p.statement(); err != nil {
			return err
		}
		p.match(RBRACE)
		if p.match(ELSE) {
			p.match(LBRACE)
			if err := p.statement(); err != nil {
				return err
			}
			p.match(RBRACE)
		}
	} else {
		return &ParseError{Message: fmt.Sprintf("Syntax error at token: %v", p.tokens[p.current])}
	}
	return nil
}

func (p *Parser) expression() {
	p.term()
	for p.match(PLUS, MINUS, GREATER_THAN) {
		p.term()
	}
}

func (p *Parser) assignment() error {
	if p.match(ID, ASSIGN) {
		p.expression()
	} else {
		return &ParseError{Message: fmt.Sprintf("Syntax error at token: %v", p.tokens[p.current])}
	}
	return nil
}

func (p *Parser) term() {
	p.factor()
	for p.match(MULTIPLY, DIVIDE) {
		p.factor()
	}
}

func (p *Parser) factor() {
	if p.match(INTEGER) || p.match(ID) {
	} else if p.match(LPAREN) {
		p.expression()
		p.match(RPAREN)
	} else if p.match(MINUS) {
		p.factor()
	} else {
		fmt.Printf("Syntax error at token: %v\n", p.tokens[p.current])
	}
}

func (p *Parser) match(expectedTypes ...TokenType) bool {
	for _, expectedType := range expectedTypes {
		if p.current < len(p.tokens) && p.tokens[p.current].Type == expectedType {
			p.current++
			return true
		}
	}
	return false
}

func main() {
  input := `
x = 3 + 5
if (x > 2) {
  y = 3
}else {
  y = 5
} `
	lexer := NewLexer(input)
	tokens := lexer.tokenize()
	fmt.Println("Tokens:", tokens)

	parser := NewParser(tokens)
	if err := parser.parse(); err != nil {
		fmt.Println("Parsing failed:", err)
	} else {
		fmt.Println("Parsing successful.")
	}
}

