package main

type Player interface {
	Play(*Game)
}

type Game struct {
	board      *Board
	dictionary *Dictionary
	letters    *LetterBag
	players    []Player
	scores     []int
	turn       int
}

func NewGame(d *Dictionary) *Game {
	g := Game{}
	g.board = &Board{}
	g.dictionary = d
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

// is every created word in the dictionary?
// XXX: make sure this function always works on its "newBoard" argument
// rather than g.board!!
//
// for each letter of the word:
// scan back, perpendicular to the word, to find the start of the other word
// if this word is only length 1, ignore it
// otherwise if it's not in the dictionary, reject
// otherwise recurse on that word
// we need the map saying which coordinates we've already checked, otherwise
// there can be cycles:
//     IN
//     NO
// we would check "IN" horizontally, then "IN" vertically, then "NO"
// horizontally, then "NO" vertically, and end up back at "IN" horizontally, etc.
func (g *Game) LegalWords(newBoard *Board, x, y int, vertical bool, checked *map[int]bool) bool {
	// walk back to find the start of the word
	for {
		nx, ny := x, y
		if vertical {
			ny--
		} else {
			nx--
		}
		if newBoard.Getchar(nx, ny) == 0 {
			break
		}
		x, y = nx, ny
	}

	// skip checks we've already done
	if (*checked)[y*15+x] {
		return true
	}
	(*checked)[y*15+x] = true

	// check this word
	word := newBoard.GetWord(x, y, vertical)

	if len(word) == 1 {
		// only 1 character long: there is no perpendicular word here
		return true
	}
	if !g.dictionary.HasWord(word) {
		return false
	}

	// check every perpendicular word
	for {
		if newBoard.Getchar(x, y) == 0 {
			break
		}
		if !g.LegalWords(newBoard, x, y, !vertical, checked) {
			return false
		}
		if vertical {
			y++
		} else {
			x++
		}
	}

	return true
}

// API for Players:

// play the given word, with legality check;
// return (score, ok)
// if not ok, you lose the turn and score 0
func (g *Game) Play(word string, x, y int, vertical bool) (int, bool) {
	defer g.NextPlayer()

	// is this play mechanically legal on this board?
	if !g.board.Legal(word, x, y, vertical) {
		return 0, false
	}

	// is this word in the dictionary?
	// XXX: we need this check separately from g.LegalWords(), otherwise any
	// single-character word would be accepted on the centre tile on the turn 1
	if !g.dictionary.HasWord(word) {
		return 0, false
	}

	// play the move on a copy of the board
	newBoard := *g.board
	score := newBoard.Play(word, x, y, vertical)

	// and check if every created word is also in the dictionary
	checked := make(map[int]bool)
	if !g.LegalWords(&newBoard, x, y, vertical, &checked) {
		return 0, false
	}

	// move accepted
	g.board = &newBoard
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
