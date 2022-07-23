package parser

import "testing"

func TestNewToken(t *testing.T) {
	newTok := newToken(NUMBER, "1")
	ttype := newTok.Ttype()
	literal := newTok.Literal()

	if ttype != NUMBER {
		t.Errorf("Ttype is not number")
	}

	if  literal != "1"{
		t.Errorf("Literal is not 1")
	}
}

func TestIsKeyWord(t *testing.T){
	if result := newToken(SELECT, "select").IsKeyWord(); !result{
		t.Errorf("select is keyword")
	}

	if result := newToken(FROM, "from").IsKeyWord(); !result{
		t.Errorf("from is keyword")
	}

	if result := newToken(WHERE, "where").IsKeyWord(); !result{
		t.Errorf("where is keyword")
	}

	if result := newToken(INSERT, "and").IsKeyWord(); !result{
		t.Errorf("insert is keyword")
	}

	if result := newToken(INTO, "into").IsKeyWord(); !result{
		t.Errorf("into is keyword")
	}

	if result := newToken(VALUES, "values").IsKeyWord(); !result{
		t.Errorf("values is keyword")
	}

	if result := newToken(DELETE, "delete").IsKeyWord(); !result{
		t.Errorf("delete is keyword")
	}

	if result := newToken(UPDATE, "update").IsKeyWord(); !result{
		t.Errorf("update is keyword")
	}

	if result := newToken(SET, "set").IsKeyWord(); !result{
		t.Errorf("set is keyword")
	}

	if result := newToken(CREATE, "create").IsKeyWord(); !result{
		t.Errorf("create is keyword")
	}

	if result := newToken(TABLE, "table").IsKeyWord(); !result{
		t.Errorf("table is keyword")
	}

	if result := newToken(INT, "int").IsKeyWord(); !result{
		t.Errorf("int is keyword")
	}

	if result := newToken(VARCHAR, "varchar").IsKeyWord(); !result{
		t.Errorf("varchar is keyword")
	}

	if result := newToken(VIEW, "view").IsKeyWord(); !result{
		t.Errorf("view is keyword")
	}

	if result := newToken(AS, "as").IsKeyWord(); !result{
		t.Errorf("as is keyword")
	}

	if result := newToken(INDEX, "index").IsKeyWord(); !result{
		t.Errorf("index is keyword")
	}

	if result := newToken(ON, "on").IsKeyWord(); !result{
		t.Errorf("on is keyword")
	}

	if result := newToken(WORD, "student").IsKeyWord(); result{
		t.Errorf("student is not keyword")
	}
}