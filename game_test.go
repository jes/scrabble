package main

import "testing"

type MockPlayer struct {
	word     []string
	x        []int
	y        []int
	vertical []bool
	turn     int
}

func TestGame(t *testing.T) {
	player1 := MockPlayer{
		word:     []string{"hello", "murder"},
		x:        []int{4, 5},
		y:        []int{7, 9},
		vertical: []bool{false, false},
	}

	player2 := MockPlayer{
		word:     []string{"loud"},
		x:        []int{6},
		y:        []int{7},
		vertical: []bool{true},
	}

	g := NewGame()
	g.AddPlayer(&player1)
	g.AddPlayer(&player2)

	checkScores(g, 0, 0, t)
	g.OneTurn()
	checkScores(g, 16, 0, t)
	g.OneTurn()
	checkScores(g, 16, 5, t)
	g.OneTurn()
	checkScores(g, 32, 5, t)
}

func checkScores(g *Game, score1, score2 int, t *testing.T) {
	if g.scores[0] != score1 || g.scores[1] != score2 {
		t.Errorf("expexted scores = %d,%d; got %d,%d", score1, score2, g.scores[0], g.scores[1])
	}
}

func (p *MockPlayer) Play(g *Game) {
	g.Play(p.word[p.turn], p.x[p.turn], p.y[p.turn], p.vertical[p.turn])
	p.turn++
}
