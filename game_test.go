package main

import (
	"fmt"
	"testing"
)

type MockPlayer struct {
	word     []string
	x        []int
	y        []int
	vertical []bool
	turn     int
	t        *testing.T
}

func TestGame(t *testing.T) {
	player1 := MockPlayer{
		word:     []string{"hello", "murder"},
		x:        []int{4, 5},
		y:        []int{7, 9},
		vertical: []bool{false, false},
		t:        t,
	}

	player2 := MockPlayer{
		word:     []string{"loud", "here"},
		x:        []int{6, 9},
		y:        []int{7, 10},
		vertical: []bool{true, false},
		t:        t,
	}

	d := NewDictionary()
	d.AddFile("dictionary")
	g := NewGame(d)
	g.AddPlayer(&player1)
	g.AddPlayer(&player2)

	checkScores(g, 0, 0, t)
	g.OneTurn()
	checkScores(g, 16, 0, t)
	g.OneTurn()
	checkScores(g, 16, 6, t)
	g.OneTurn()
	checkScores(g, 33, 6, t)
	g.OneTurn()
	checkScores(g, 33, 29, t)
}

func checkScores(g *Game, score1, score2 int, t *testing.T) {
	if g.scores[0] != score1 || g.scores[1] != score2 {
		fmt.Println(g.board)
		t.Errorf("expected scores = %d,%d; got %d,%d", score1, score2, g.scores[0], g.scores[1])
	}
}

func (p *MockPlayer) Play(g *Game) {
	if !g.Legal(p.word[p.turn], p.x[p.turn], p.y[p.turn], p.vertical[p.turn]) {
		p.t.Errorf("'%s' at %d,%d is illegal; expected it to be legal", p.word[p.turn], p.x[p.turn], p.y[p.turn])
	}

	g.Play(p.word[p.turn], p.x[p.turn], p.y[p.turn], p.vertical[p.turn])
	p.turn++
}
