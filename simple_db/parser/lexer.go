package parser

import (
	"strings"
	"github.com/pkg/errors"
	"fmt"
	"strconv"
)

const (
	NULL_NUMBER = 0
	SINGLE_QUOTATION_NUMBER = 39
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

func (l *lexer) matchIntConstant() bool{
	return NUMBER == l.tok.Ttype()
}

func (l *lexer) matchStringConstant() bool {
	return SINGLE_QUOTATION == l.tok.Ttype()
}

func (l *lexer) matchKeyword(w string) bool{
	return WORD == l.tok.Ttype() && l.tok.Literal() == w
}

func (l *lexer) matchID() bool{
	return WORD == l.tok.Ttype() && !l.tok.IsKeyWord()
}

func (l *lexer) eatDelim(d TokenType) (error){
	if !l.matchDelim(d){
		return errors.New(fmt.Sprintf("BadSyntaxException"))
	}

	l.nextToken();

	return nil
}

func (l *lexer) eatIntConstant() (int, error){
	if !l.matchIntConstant(){
		return 0, errors.New(fmt.Sprintf("BadSyntaxException"))
	}

	literal := l.tok.Literal()
	i, _ := strconv.Atoi(literal)
	l.nextToken();

	return  i, nil
}

func (l *lexer) eatStringConstant() (string, error){
	if !l.matchStringConstant(){
		return "", errors.New(fmt.Sprintf("BadSyntaxException"))
	}

	literal := l.tok.Literal()
	l.nextToken();

	return literal, nil
}

func (l *lexer) eatKeyword(w string) error{
	if (!l.matchKeyword(w)){
		return errors.New(fmt.Sprintf("BadSyntaxException"))
	}

	l.nextToken();

	return nil
}

func (l *lexer) eatID() (string, error){
	if (!l.matchID()){
		return "", errors.New(fmt.Sprintf("BadSyntaxException"))
	}

	literal := l.tok.Literal()
	l.nextToken();
	return literal, nil
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

func (l *lexer) nextToken(){

	l.skipWhitespace()

	switch l.ch{
	case NULL_NUMBER:
		l.tok = newToken(EOF, "")
	case SINGLE_QUOTATION_NUMBER:
		l.tok = newToken(SINGLE_QUOTATION, string(l.ch))
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