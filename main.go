package main

import (
	"Lexer/Lexer"
	"fmt"
)

func main() {
	input := `
let ten = 10;
let sqr = fn(x) { x * x; }
let result = sqr(ten);

if (res > ten) {
	return true;
} else {
	return false;
}

if (10 == 10) {
	return true;
}

if (10 != 9) {
	return true;
}
`
	var lexer = Lexer.NewTokenizer(input)

	for {
		t := lexer.GetNextToken()
		if t.TokenType == Lexer.TokenEof {
			break
		}

		fmt.Println(t)
	}
}
