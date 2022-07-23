package parser

type TokenType string

type token struct{
	ttype TokenType
	literal string
}

func newToken(ttype TokenType, literal string) *token{
	return &token{ttype: ttype, literal: literal}
}

func (t *token)IsKeyWord()bool{
	if _, ok := keywords[t.literal]; ok{
		return true
	}

	return false
}

func (t *token)Ttype()TokenType{
	return t.ttype
}

func (t *token)Literal()string{
	return t.literal
}

const (
	EOF = "EOF"
	ILLEGAL = "ILLEGAL"
	NUMBER = "number"
	WORD = "word"
	SELECT = "select"
	FROM = "from"
	WHERE = "where"
	AND = "and"
	INSERT = "insert"
	INTO = "into"
	VALUES = "values"
	DELETE = "delete"
	UPDATE = "update"
	SET = "set"
	CREATE = "create"
	TABLE = "table"
	INT = "int"
	VARCHAR = "varchar"
	VIEW = "view"
	AS = "as"
	INDEX = "index"
	ON = "on"
	SINGLE_QUOTATION = "SINGLE_QUOTATION"
)

var keywords = map[string]TokenType {
	"select": SELECT,
	"from": FROM,
	"where": WHERE,
	"and": AND,
	"insert": INSERT,
	"into": INTO,
	"values": VALUES,
	"delete": DELETE,
	"update": UPDATE,
	"set": SET,
	"create": CREATE,
	"table": TABLE,
	"int": INT,
	"varchar": VARCHAR,
	"view": VIEW,
	"as": AS,
	"index": INDEX,
	"on": ON,
}