package Lexer

import "testing"

func TestLexerGetNextToken(t *testing.T) {
	var input = `=+(){},;`
	var expecting = []Token{
		{tokenAssign, "="},
		{tokenPlus, "+"},
		{tokenLParen, "("},
		{tokenRParen, ")"},
		{tokenLCurly, "{"},
		{tokenRCurly, "}"},
		{tokenComma, ","},
		{tokenSemicolon, ";"},
	}

	var lexer = NewTokenizer(input)

	for _, token := range expecting {
		var nextToken = lexer.GetNextToken()
		if nextToken.TokenType != token.TokenType {
			t.Errorf("GetNextToken expected '%s', got '%v'", token, nextToken)
		}
	}
}

func TestLexerParseString(t *testing.T) {
	const input = `
let five = 5;
let ten = 10;
let add = fn(x,y) { x + y; };
let res = add(five, ten);
`

	var lexer = NewTokenizer(input)

	tokens := []Token{

		CreateToken(tokenLet, "let"),
		CreateToken(tokenIdent, "five"),
		CreateToken(tokenAssign, "="),
		CreateToken(tokenInt, "5"),
		CreateToken(tokenSemicolon, ";"),

		CreateToken(tokenLet, "let"),
		CreateToken(tokenIdent, "ten"),
		CreateToken(tokenAssign, "="),
		CreateToken(tokenInt, "10"),
		CreateToken(tokenSemicolon, ";"),

		CreateToken(tokenLet, "let"),
		CreateToken(tokenIdent, "add"),
		CreateToken(tokenAssign, "="),
		CreateToken(tokenFunction, "fn"),
		CreateToken(tokenLParen, "("),
		CreateToken(tokenIdent, "x"),
		CreateToken(tokenComma, ","),
		CreateToken(tokenIdent, "y"),
		CreateToken(tokenRParen, ")"),
		CreateToken(tokenLCurly, "{"),
		CreateToken(tokenIdent, "x"),
		CreateToken(tokenPlus, "+"),
		CreateToken(tokenIdent, "y"),
		CreateToken(tokenSemicolon, ";"),
		CreateToken(tokenRCurly, "}"),
		CreateToken(tokenSemicolon, ";"),

		CreateToken(tokenLet, "let"),
		CreateToken(tokenIdent, "result"),
		CreateToken(tokenAssign, "="),
		CreateToken(tokenIdent, "add"),
		CreateToken(tokenLParen, "("),
		CreateToken(tokenIdent, "five"),
		CreateToken(tokenComma, ","),
		CreateToken(tokenIdent, "ten"),
		CreateToken(tokenRParen, ")"),
		CreateToken(tokenSemicolon, ";"),
		CreateToken(TokenEof, "eof"),
	}

	for _, token := range tokens {
		var nextToken = lexer.GetNextToken()
		if nextToken.TokenType != token.TokenType {
			t.Errorf("Token expected '%s', got '%v'", token.TokenType, nextToken)
		}
	}
}
