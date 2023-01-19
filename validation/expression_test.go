package validation

import (
	"LL/table"
	"testing"
)

var testTable = table.IntiTable("../table.csv")

func TestIsLineValid(t *testing.T) {
	line := "-8*--(---3+a*b*--8*(a+--3))"
	t.Log(line)

	assertTrue(t, IsLineValid(line, testTable))

	line = "-8*-)(-3+a*b*-8*(a+-3))"
	t.Log(line)

	assertFalse(t, IsLineValid(line, testTable))

	line = "-8*-(-3++a*b*-8*(a+++-3))"
	t.Log(line)

	assertFalse(t, IsLineValid(line, testTable))

	line = "-8**-(-3+a***b****-8**(a+-3))"
	t.Log(line)

	assertFalse(t, IsLineValid(line, testTable))

	line = "-*-(-3+a*b*-8*(a+-3))"
	t.Log(line)

	assertFalse(t, IsLineValid(line, testTable))

	line = ""
	t.Log("empty line")

	assertFalse(t, IsLineValid(line, testTable))

	line = "-8*-(-3+a*b*-8*(a+-3))"
	t.Log(line)

	assertTrue(t, IsLineValid(line, testTable))

	line = "-8*-(-3+(a*b*-8)*(a+-3))"
	t.Log(line)

	assertTrue(t, IsLineValid(line, testTable))
}

func assertTrue(t *testing.T, val bool) {
	if !val {
		t.Fatalf("should be true")
	}
}

func assertFalse(t *testing.T, val bool) {
	if val {
		t.Fatalf("should be false")
	}
}
