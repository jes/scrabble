package main

type Player interface {
	Play(*Game)
}

type Game struct {
	board   *Board
	letters *LetterBag
	players []Player
	scores  []int
	turn    int
}

func NewGame(d *Dictionary) *Game {
	g := Game{}
	board := NewBoard(d)
	g.board = &board
	g.letters = NewLetterBag()

	return &g
}

func (g *Game) AddPlayer(p Player) {
	g.players = append(g.players, p)
	g.scores = append(g.scores, 0)
}

func (g *Game) OneTurn() {
	g.players[g.turn].Play(g)
}

func (g *Game) NextPlayer() {
	g.turn = (g.turn + 1) % len(g.players)
}

// API for Players:

func (g *Game) Legal(word string, x, y int, vertical bool) bool {
	return g.board.Legal(word, x, y, vertical)
}

// play the given word, with legality check;
// return (score, ok)
// if not ok, you lose the turn and score 0
func (g *Game) Play(word string, x, y int, vertical bool) (int, bool) {
	defer g.NextPlayer()

	// is this play mechanically legal on this board?
	if !g.board.Legal(word, x, y, vertical) {
		return 0, false
	}

	// play the move
	score := g.board.Play(word, x, y, vertical)
	g.scores[g.turn] += score

	return score, true
}

// Swap some letters with letters from the bag (pass in
// the letters to discard, return the new letters)
// if there aren't enough letters left in the bag
// then only the amount that are left will be swapped
func (g *Game) SwapLetters(letters []byte) []byte {
	newLetters := make([]byte, 0)

	rebag := make([]byte, 0)

	for i := range letters {
		c := letters[i]
		if !g.letters.BagEmpty() {
			// make sure to put the old letter back in the bag
			rebag = append(rebag, c)

			// take next letter from bag
			c = g.letters.NextLetter()
		}
		newLetters = append(newLetters, c)
	}

	// put the swapped letters back in the bag
	for _, c := range rebag {
		g.letters.AppendLetter(c)
	}

	g.NextPlayer()

	return newLetters
}

// Skip the turn without swapping any letters
func (g *Game) SkipTurn() {
	g.NextPlayer()
}
