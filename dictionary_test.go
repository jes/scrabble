package main

import "testing"

func TestDictionary(t *testing.T) {
	d := NewDictionary()
	if d.HasWord("foo") {
		t.Errorf("'foo' shouldn't exist yet, but does")
	}
	d.AddWord("foo")
	if !d.HasWord("foo") {
		t.Errorf("'foo' should exist now, but doesn't")
	}

	if d.HasWord("abacus") {
		t.Errorf("'abacus' shouldn't exist yet, but does")
	}
	d.AddFile("dictionary")
	if !d.HasWord("abacus") {
		t.Errorf("'abacus' should exist now, but doesn't")
	}

	if !d.HasWord("gazelle") {
		t.Errorf("'gazelle' should exist, but doesn't")
	}
	if !d.HasWord("gazELLe") {
		t.Errorf("'gazELLe' should exist, but doesn't (dictionary should be case insensitive)")
	}
}
