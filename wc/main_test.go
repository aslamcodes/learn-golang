package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b:= bytes.NewBufferString("word1 word2")
	exp := 2
	res := count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b:= bytes.NewBufferString("word1 word2")
	exp := 1
	res := count(b, true, false)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b:= bytes.NewBufferString("word1 🥲")
	exp := 10
	res := count(b, false, true)
	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}
