package main

type Player interface {
	Play(*Game)
}

type Game struct {
	board   *Board
	letters *LetterBag
	players []*Player
	scores  []int
}

func NewGame() Game {
	g := Game{}
	g.board = &Board{}
	g.letters = NewLetterBag()

	return g
}

func (g *Game) AddPlayer(p *Player) {
	g.players = append(g.players, p)
	g.scores = append(g.scores, 0)
}
