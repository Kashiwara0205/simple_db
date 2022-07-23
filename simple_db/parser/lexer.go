package parser

type lexer struct{
	input string
	position int
	readPosition int
	ch byte
}

func NewLexer(input string) *lexer{
	l := &lexer{input: input}

	l.readChar()

	return l
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