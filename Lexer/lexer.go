package Lexer

type TokenType string

const (
	Illegal    TokenType = "ILLEGAL"
	TokenEof   TokenType = "EOF"
	tokenIdent TokenType = "IDENT"

	tokenInt    TokenType = "INT"
	tokenAssign TokenType = "="
	tokenBang   TokenType = "!"

	tokenPlus     TokenType = "+"
	tokenMinus    TokenType = "-"
	tokenMultiply TokenType = "*"
	tokenDivide   TokenType = "/"

	tokenComma     TokenType = ","
	tokenSemicolon TokenType = ";"

	tokenLParen TokenType = "("
	tokenRParen TokenType = ")"
	tokenLCurly TokenType = "{"
	tokenRCurly TokenType = "}"
	tokenLAngle TokenType = "<"
	tokenRAngle TokenType = ">"

	tokenFunction TokenType = "FUNCTION"
	tokenLet      TokenType = "LET"

	tokenTrue   = "TRUE"
	tokenFalse  = "FALSE"
	tokenIf     = "IF"
	tokenElse   = "ELSE"
	tokenReturn = "RETURN"
)

type Token struct {
	TokenType  TokenType
	StrLiteral string
}

func CreateToken(tokenType TokenType, strLiteral string) Token {
	return Token{tokenType, strLiteral}
}

const _0 = int('0')
const _9 = int('9')

const a = int('a')
const z = int('z')

const A = int('A')
const Z = int('Z')

const __ = int('_')

func isLetter(chr rune) bool {
	var char = int(chr)
	return a <= char && z >= char ||
		A <= char && z >= char ||
		char == __
}

func isNumber(chr rune) bool {
	var char = int(chr)
	return _0 <= char && _9 >= char
}

var Keyword = map[string]Token{
	"fn":     CreateToken(tokenFunction, "fn"),
	"let":    CreateToken(tokenLet, "let"),
	"true":   CreateToken(tokenTrue, "true"),
	"false":  CreateToken(tokenFalse, "false"),
	"if":     CreateToken(tokenIf, "if"),
	"else":   CreateToken(tokenElse, "else"),
	"return": CreateToken(tokenReturn, "return"),
}

type Tokenizer struct {
	position     int
	readPosition int
	chr          rune
	input        string
}

func NewTokenizer(input string) *Tokenizer {
	var tokenizer = Tokenizer{
		position:     0,
		readPosition: 0,
		input:        input,
	}
	tokenizer.readChar()
	return &tokenizer
}

func (tokenizer *Tokenizer) GetNextToken() Token {
	tokenizer.skipWhitespace()

	var tok Token
	var tokNil bool = true

	switch tokenizer.chr {
	case '{':
		tok = CreateToken(tokenLCurly, string(tokenizer.chr))
		tokNil = false
	case '}':
		tok = CreateToken(tokenRCurly, string(tokenizer.chr))
		tokNil = false
	case '(':
		CreateToken(tokenLParen, string(tokenizer.chr))
		tokNil = false
	case ')':
		tok = CreateToken(tokenRParen, string(tokenizer.chr))
		tokNil = false
	case '>':
		tok = CreateToken(tokenRAngle, string(tokenizer.chr))
		tokNil = false
	case '<':
		tok = CreateToken(tokenLAngle, string(tokenizer.chr))
		tokNil = false
	case ';':
		tok = CreateToken(tokenSemicolon, string(tokenizer.chr))
		tokNil = false
	case ',':
		tok = CreateToken(tokenComma, string(tokenizer.chr))
		tokNil = false
	case '+':
		tok = CreateToken(tokenPlus, string(tokenizer.chr))
		tokNil = false
	case '*':
		tok = CreateToken(tokenMultiply, string(tokenizer.chr))
		tokNil = false
	case '/':
		tok = CreateToken(tokenDivide, string(tokenizer.chr))
		tokNil = false
	case '=':
		tok = CreateToken(tokenAssign, string(tokenizer.chr))
		tokNil = false
	case '!':
		tok = CreateToken(tokenBang, string(tokenizer.chr))
		tokNil = false
	case '\x00':
		tok = CreateToken(TokenEof, "eof")
		tokNil = false
	}

	if isLetter(tokenizer.chr) {
		var ident = tokenizer.readIdentifier()
		var keyword, exists = Keyword[ident]
		if exists {
			return keyword
		} else {
			return CreateToken(tokenIdent, ident)
		}
	} else if isNumber(tokenizer.chr) {
		return CreateToken(tokenInt, tokenizer.readInt())
	} else if tokNil {
		return CreateToken(Illegal, string(tokenizer.chr))
	}

	tokenizer.readChar()
	return tok
}

func (tokenizer *Tokenizer) readChar() {
	if tokenizer.readPosition >= len(tokenizer.input) {
		tokenizer.chr = '\x00'
	} else {
		tokenizer.chr = rune(tokenizer.input[tokenizer.readPosition])
	}

	tokenizer.position = tokenizer.readPosition
	tokenizer.readPosition++
}

func (tokenizer *Tokenizer) readIdentifier() string {
	var position = tokenizer.position
	for isLetter(tokenizer.chr) {
		tokenizer.readChar()
	}

	return tokenizer.input[position:tokenizer.position]
}

func (tokenizer *Tokenizer) readInt() string {
	var position = tokenizer.position
	for isNumber(tokenizer.chr) {
		tokenizer.readChar()
	}

	return tokenizer.input[position:tokenizer.position]
}

func (tokenizer *Tokenizer) skipWhitespace() {
	for tokenizer.chr == ' ' || tokenizer.chr == '\t' || tokenizer.chr == '\n' || tokenizer.chr == '\r' {
		tokenizer.readChar()
	}
}
