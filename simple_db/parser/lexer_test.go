package parser

import "testing"

func TestNewLexer(t *testing.T) {
	lexer := NewLexer("select a1")

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
	lexer := NewLexer("select")

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