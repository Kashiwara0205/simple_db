package parser

import "testing"

func TestNewLexer(t *testing.T) {
	lexer := newLexer("select a1")

	if lexer.input != "select a1" {
		t.Errorf("lexer.input is not select a1")
	}

	if lexer.position != 0 {
		t.Errorf("lexer.position is not 0")
	}

	if lexer.readPosition != 1 {
		t.Errorf("lexer.readPosition is not 1")
	}

	if lexer.ch != 's' {
		t.Errorf("lexer.ch is not s")
	}
}

func TestReadChar(t *testing.T) {
	var lexer = newLexer("select")

	if lexer.ch != 's' {
		t.Errorf("lexer.ch is not s")
	}

	if lexer.readChar(); lexer.ch != 'e' {
		t.Errorf("lexer.ch is not e")
	}

	if lexer.readChar(); lexer.ch != 'l' {
		t.Errorf("lexer.ch is not l")
	}

	if lexer.readChar(); lexer.ch != 'e' {
		t.Errorf("lexer.ch is not e")
	}

	if lexer.readChar(); lexer.ch != 'c' {
		t.Errorf("lexer.ch is not c")
	}

	if lexer.readChar(); lexer.ch != 't' {
		t.Errorf("lexer.ch is not t")
	}

	if lexer.readChar(); lexer.ch != 0 {
		t.Errorf("lexer.ch is not 0")
	}

	lexer = newLexer("SELECT")

	if lexer.ch != 's' {
		t.Errorf("lexer.ch is not s")
	}

	if lexer.readChar(); lexer.ch != 'e' {
		t.Errorf("lexer.ch is not e")
	}

	if lexer.readChar(); lexer.ch != 'l' {
		t.Errorf("lexer.ch is not l")
	}

	if lexer.readChar(); lexer.ch != 'e' {
		t.Errorf("lexer.ch is not e")
	}

	if lexer.readChar(); lexer.ch != 'c' {
		t.Errorf("lexer.ch is not c")
	}

	if lexer.readChar(); lexer.ch != 't' {
		t.Errorf("lexer.ch is not t")
	}

	if lexer.readChar(); lexer.ch != 0 {
		t.Errorf("lexer.ch is not 0")
	}
}

func TestSkipWhiteSpace(t *testing.T) {
	var lexer = newLexer("a b")
	lexer.readChar();
	lexer.skipWhitespace();

	if  lexer.ch != 'b'{
		t.Errorf("lexer.ch is not b")
	}

	lexer = newLexer("a\nb")
	lexer.readChar();
	lexer.skipWhitespace();

	if  lexer.ch != 'b'{
		t.Errorf("lexer.ch is not b")
	}

	lexer = newLexer("a\tb")
	lexer.readChar();
	lexer.skipWhitespace();

	if  lexer.ch != 'b'{
		t.Errorf("lexer.ch is not b")
	}

	lexer = newLexer("a\rb")
	lexer.readChar();
	lexer.skipWhitespace();

	if  lexer.ch != 'b'{
		t.Errorf("lexer.ch is not b")
	}
}

func TestIsLetter(t *testing.T) {
	if  !isLetter('a'){
		t.Errorf("a is letter")
	}

	if  !isLetter('z'){
		t.Errorf("z is letter")
	}

	if  !isLetter('A'){
		t.Errorf("A is letter")
	}

	if  !isLetter('Z'){
		t.Errorf("Z is letter")
	}

	if  !isLetter('_'){
		t.Errorf("_ is letter")
	}

	if  isLetter('2'){
		t.Errorf("2 is not letter")
	}
}

func TestIsDigit(t *testing.T) {
	if  !isDigit('0'){
		t.Errorf("0 is digit")
	}

	if  !isDigit('9'){
		t.Errorf("9 is digit")
	}

	if  isDigit('a'){
		t.Errorf("a is not digit")
	}
}

func TestReadNumber(t *testing.T){
	lexer := newLexer("222 hoge")
	if number := lexer.readNumber(); number != "222"{
		t.Errorf("number is not 222")
	}
}

func TestReadWord(t *testing.T){
	lexer := newLexer("select 222")
	if word := lexer.readWord(); word != "select"{
		t.Errorf("word is not select")
	}
}

func TestNextToken(t *testing.T){
	lexer := newLexer("select a1")
	lexer.NextToken()

	if lexer.tok.Literal() != "select"{
		t.Errorf("tok.Literal is not select")
	}

	if lexer.tok.Ttype() != WORD{
		t.Errorf("tok.Ttype is not WORD")
	}

	lexer.NextToken()

	if lexer.tok.Literal() != "a"{
		t.Errorf("tok.Literal is not a")
	}

	if lexer.tok.Ttype() != WORD{
		t.Errorf("tok.Ttype is not WORD")
	}

	lexer.NextToken()

	if lexer.tok.Literal() != "1"{
		t.Errorf("tok.Literal is not 1")
	}

	if lexer.tok.Ttype() != NUMBER{
		t.Errorf("tok.Ttype is not NUMBER")
	}

	lexer.NextToken()

	if lexer.tok.Literal() != ""{
		t.Errorf("tok.Literal is not ''")
	}

	if lexer.tok.Ttype() != EOF{
		t.Errorf("tok.Ttype is not EOF")
	}
}

func TestMatchDelim(t *testing.T){
	var lexer = newLexer("select")
	lexer.NextToken()

	if !lexer.matchDelim(WORD){
		t.Errorf("tok.Ttype is not WORD")
	}

	lexer = newLexer("100")
	lexer.NextToken()

	if !lexer.matchDelim(NUMBER){
		t.Errorf("tok.Ttype is not NUMBER")
	}
}