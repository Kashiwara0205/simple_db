package parser

import (
	"strings"
)

type lexer struct{
	input string
	position int
	readPosition int
	ch byte
	tok *token
}

func newLexer(input string) *lexer{
	l := &lexer{input: strings.ToLower(input)}

	l.readChar()

	return l
}

func (l *lexer) matchDelim(d TokenType) bool{
	return d == l.tok.Ttype()
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *lexer) skipWhitespace(){
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r'{
		l.readChar()
	}
}

func (l *lexer) NextToken(){

	l.skipWhitespace()

	switch l.ch{
	case 0:
		l.tok = newToken(EOF, "")
	default:
		if isLetter(l.ch){
			l.tok = newToken(WORD, l.readWord())
		}else if isDigit(l.ch){
			l.tok = newToken(NUMBER, l.readNumber())
		}else{
			l.tok = newToken(ILLEGAL, string(l.ch))
		}
	}
}

func (l *lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch){
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *lexer) readWord() string{
	position := l.position

	for isLetter(l.ch){
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool{
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool{
	return '0' <= ch && ch <= '9'
}