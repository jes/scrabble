package main

import (
	"fmt"
	"math/rand"
	"os"
)

type RandomAI struct {
	letters     *LetterBag
	nwords      int
	maxattempts int
}

func NewRandomAI(nwords, maxattempts int) *RandomAI {
	ai := RandomAI{}
	ai.letters = NewLetterBag()
	ai.letters.Empty()
	ai.nwords = nwords
	ai.maxattempts = maxattempts
	return &ai
}

func (ai *RandomAI) Play(g *Game) {
	// draw before the turn so that we have some letters on
	// the first turn (this will be a no-op on every subsequent
	// turn because we also draw after we make the move)
	ai.DrawLetters(g)

	bestword, bestx, besty, bestvertical, bestscore := "", 0, 0, false, -1

	var bestletterbag []byte

	for i := 0; i < ai.nwords; i++ {
		word, x, y, vertical, letterbag := ai.AnyMove(g)
		score := g.board.Score(word, x, y, vertical)
		if score > bestscore {
			bestword, bestx, besty, bestvertical, bestletterbag, bestscore = word, x, y, vertical, letterbag, score
		}
	}

	if bestword == "" {
		ai.letters.letters = g.SwapLetters(ai.letters.letters)
	} else {
		_, ok := g.Play(bestword, bestx, besty, bestvertical)
		if !ok {
			fmt.Fprintf(os.Stderr, "%s at %d,%d(%v): illegal move", bestword, bestx, besty, bestvertical)
		}

		ai.letters.letters = bestletterbag
	}

	ai.DrawLetters(g)
}

func (ai *RandomAI) AnyMove(g *Game) (string, int, int, bool, []byte) {
	originalLetterBag := make([]byte, len(ai.letters.letters))
	copy(originalLetterBag, ai.letters.letters)

	for i := 0; i < ai.maxattempts; i++ {
		nletters := rand.Intn(ai.letters.BagSize()-1) + 1
		x := rand.Intn(15)
		y := rand.Intn(15)
		vertical := rand.Intn(1) == 0

		// find the start point of the word
		if vertical {
			for ; y > 0 && g.board.Getchar(x, y) != 0; y-- {
			}
		} else {
			for ; x > 0 && g.board.Getchar(x, y) != 0; x-- {
			}
		}

		// build the word
		dx, dy := 1, 0
		if vertical {
			dx, dy = 0, 1
		}
		n := 0
		word := ""
		for x < 15 && y < 15 {
			if g.board.Getchar(x, y) == 0 {
				if n >= nletters {
					break
				}
				word += string(ai.letters.NextLetter())
				n++
			} else {
				word += string(g.board.Getchar(x, y))
			}
			x += dx
			y += dy
		}

		lettersRemaining := ai.letters.letters
		ai.letters.letters = originalLetterBag

		// if the move is legal, we're done!
		if g.board.Legal(word, x, y, vertical) {
			return word, x, y, vertical, lettersRemaining
		}

		// otherwise try again
	}

	return "", 0, 0, false, nil
}

func (ai *RandomAI) DrawLetters(g *Game) {
	for ai.letters.BagSize() < 7 {
		ch, ok := g.GetLetter()
		if !ok {
			break
		}
		ai.letters.AppendLetter(ch)
	}
}
