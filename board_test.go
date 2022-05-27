package main

import "testing"

func TestChars(t *testing.T) {
	b := Board{}
	b.Putchar('x', 0, 0)
	b.Putchar('x', -1, -1)
	b.Putchar('o', 10, 14)
	b.Putchar('N', 10, 14)
	b.Putchar('N', 15, 14)
	checkCell(&b, 0, 0, 'x', t)
	checkCell(&b, -1, -1, 0, t)
	checkCell(&b, 10, 14, 'N', t)
	checkCell(&b, 15, 14, 0, t)
}

func checkCell(b *Board, x, y int, c byte, t *testing.T) {
	if b.Getchar(x, y) != c {
		t.Errorf("expected %c at %d,%d; got %c", c, x, y, b.Getchar(x, y))
	}
}

func TestScore(t *testing.T) {
	b := Board{}
	if b.Score("hello", 4, 7, false) != 16 {
		t.Errorf("expected 16 points for 'hello' at 4,7; got %d", b.Score("hello", 4, 7, false))
	}
	checkPlay(&b, "hello", 4, 7, false, 16, t)
	checkPlay(&b, "loud", 6, 7, true, 5, t)
	checkPlay(&b, "murder", 5, 9, false, 16, t)
}

func checkPlay(b *Board, word string, x, y int, vert bool, wantScore int, t *testing.T) {
	score := b.Play(word, x, y, vert)
	if score != wantScore {
		t.Errorf("expected %d points for '%s' at %d,%d; got %d", wantScore, word, x, y, score)
	}
}

func TestLegal(t *testing.T) {
	b := Board{}
	checkLegal(&b, "hello", 4, 6, false, false, t)
	checkLegal(&b, "hello", 4, 7, false, true, t)
	checkLegal(&b, "loud", 7, 7, false, false, t)
	checkLegal(&b, "loud", 6, 7, true, true, t)
	checkLegal(&b, "murder", 5, 9, false, true, t)
	checkLegal(&b, "foo", 0, 0, false, false, t)
}

func checkLegal(b *Board, word string, x, y int, vert bool, wantLegal bool, t *testing.T) {
	legal := b.Legal(word, x, y, vert)
	if legal != wantLegal {
		legalWord := map[bool]string{
			true:  "legal",
			false: "illegal",
		}
		t.Errorf("expected %s for '%s' at %d,%d; got %s", legalWord[wantLegal], word, x, y, legalWord[legal])
	}
	if legal {
		b.Play(word, x, y, vert)
	}
}
