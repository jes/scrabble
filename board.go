package main

type Board struct {
	cell       [225]byte
	dictionary *Dictionary
}

var wordMultipleMap = map[int]int{
	0: 3, 7: 3, 14: 3,
	16: 2, 28: 2,
	32: 2, 42: 2,
	48: 2, 56: 2,
	64: 2, 70: 2,
	105: 3, 112: 2, 119: 3,
	154: 2, 160: 2,
	168: 2, 176: 2,
	182: 2, 192: 2,
	196: 2, 208: 2,
	210: 3, 217: 3, 224: 3,
}
var letterMultipleMap = map[int]int{
	3: 2, 11: 2,
	20: 3, 24: 3,
	36: 2, 38: 2,
	45: 2, 52: 2, 59: 2,
	76: 3, 80: 3, 84: 3, 88: 3,
	92: 2, 96: 2, 98: 2, 102: 2,
	108: 2, 116: 2,
	122: 2, 126: 2, 128: 2, 132: 2,
	136: 3, 140: 3, 144: 3, 148: 3,
	165: 2, 172: 2, 179: 2,
	186: 2, 188: 2,
	200: 2, 204: 2,
	213: 2, 221: 2,
}
var letterScore = map[byte]int{
	'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1,
	'f': 4, 'g': 2, 'h': 4, 'i': 1, 'j': 8,
	'k': 5, 'l': 1, 'm': 3, 'n': 1, 'o': 1,
	'p': 3, 'q': 10, 'r': 1, 's': 1, 't': 1,
	'u': 1, 'v': 4, 'w': 4, 'x': 8, 'y': 4,
	'z': 10,
}

func NewBoard(d *Dictionary) Board {
	return Board{dictionary: d}
}

// return the score for playing the given word at the given position, with no legality check;
// words should be all lowercase, except for blank tiles which should be uppercase
func (b *Board) Score(word string, x, y int, vertical bool) int {
	score := 0

	// play the word on a copy of the board
	newBoard := *b
	newBoard.PutWord(word, x, y, vertical)

	b.ScanWords(&newBoard, x, y, vertical, func(word string, x, y int, vertical bool) {
		score += b.OneWordScore(word, x, y, vertical)
	}, true)

	return score
}

// score just 1 word, without regard for adjacent words
func (b *Board) OneWordScore(word string, x, y int, vertical bool) int {
	score := 0

	wordFactor := 1

	for i := 0; i < len(word); i++ {
		if b.Getchar(x, y) == 0 {
			score += b.LetterMultiple(x, y) * letterScore[word[i]]
			wordFactor *= b.WordMultiple(x, y)
		} else {
			score += letterScore[word[i]]
		}
		if vertical {
			y++
		} else {
			x++
		}
	}

	return score * wordFactor
}

// play the word with no legality check, return its score;
// words should be all lowercase, except for blank tiles which should be uppercase
func (b *Board) Play(word string, x, y int, vertical bool) int {
	score := b.Score(word, x, y, vertical)

	b.PutWord(word, x, y, vertical)

	return score
}

// is every created word in the dictionary?
func (b *Board) LegalWords(newBoard *Board, x, y int, vertical bool) bool {
	if !b.dictionary.HasWord(newBoard.GetWord(x, y, vertical)) {
		return false
	}

	allLegal := true
	b.ScanWords(newBoard, x, y, vertical, func(word string, x, y int, vertical bool) {
		allLegal = allLegal && b.dictionary.HasWord(word)
	}, true)
	return allLegal
}

// find the new words created in newBoard starting from x,y, and
// pass them to cb()
//
// scan backwards to find the start of the word
// if the word is only length 1, ignore it
// if the word contains no newly-played characters, ignore it
// call cb()
// if "recurse", then recurse for each perpendicular word
func (b *Board) ScanWords(newBoard *Board, x, y int, vertical bool, cb func(string, int, int, bool), recurse bool) {
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

	// check that this word is new
	word := newBoard.GetWord(x, y, vertical)
	existingWord := b.GetWord(x, y, vertical)
	if word == existingWord {
		return
	}

	if len(word) == 1 {
		// only 1 character long: there is no perpendicular word here
		return
	}

	// pass word to callback
	cb(word, x, y, vertical)

	// recurse for every perpendicular word
	if recurse {
		for {
			if newBoard.Getchar(x, y) == 0 {
				break
			}
			b.ScanWords(newBoard, x, y, !vertical, cb, false)
			if vertical {
				y++
			} else {
				x++
			}
		}
	}
}

// return true if the given word:
// - fits within the board confines,
// - and is at least 2 characters long,
// - and doesn't conflict with any existing letters,
// - and touches at least one existing letter or places the centre tile,
// - and places at least one new tile,
// - and there is a blank (or edge) immediately before and immediately after the word,
// - and this word and every adjacent word is in the dictionary
func (b *Board) Legal(word string, x, y int, vertical bool) bool {
	endx := x
	endy := y

	if vertical {
		endy += len(word)
	} else {
		endx += len(word)
	}

	// bounds check
	if x < 0 || endx >= 15 || y < 0 || endy >= 15 {
		return false
	}

	// disallow length-1 words, otherwise you can gain
	// spurious points, for example turning "head"
	// into "ahead" by claiming that you're playing the vertical
	// word "a" and scoring an additional horizontal word "ahead"
	if len(word) <= 1 {
		return false
	}

	// must be a blank (or edge) immediately before and immediately after
	// e.g. you can't pretend to play "FOO" by sticking it on the end of "ELF",
	// because that is "ELFOO"; you need to declare the true word played
	xpre, ypre := x, y
	xpost, ypost := endx, endy
	if vertical {
		ypre--
		ypost++
	} else {
		xpre--
		xpost++
	}
	if b.Getchar(xpre, ypre) != 0 || b.Getchar(xpost, ypost) != 0 {
		return false
	}

	gotNewTile := false
	gotCentreTile := false
	gotAdjacentTile := false

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	px, py := x, y

	for i := 0; i < len(word); i++ {
		c := b.Getchar(px, py)
		// reject non-matching chars
		if c != 0 && lc(c) != lc(word[i]) {
			return false
		}
		if c == 0 {
			gotNewTile = true
			if px == 7 && py == 7 {
				gotCentreTile = true
			}
		}

		for d := range dx {
			if b.Getchar(px+dx[d], py+dy[d]) != 0 {
				gotAdjacentTile = true
			}
		}

		if vertical {
			py++
		} else {
			px++
		}
	}

	// need to play at least 1 tile
	if !gotNewTile {
		return false
	}

	// need to play adjacent to an existing word, or on the centre tile
	if !gotCentreTile && !gotAdjacentTile {
		return false
	}

	// play the move on a copy of the board and work out whether
	// this word and all of the adjacent words are in the dictionary
	if b.dictionary != nil {
		newBoard := *b
		newBoard.Play(word, x, y, vertical)
		if !b.LegalWords(&newBoard, x, y, vertical) {
			return false
		}
	}

	return true
}

func (b *Board) Putchar(char byte, x, y int) {
	if x < 0 || x >= 15 || y < 0 || y >= 15 {
		return
	}
	b.cell[y*15+x] = char
}

func (b *Board) Getchar(x, y int) byte {
	if x < 0 || x >= 15 || y < 0 || y >= 15 {
		return 0
	}
	return b.cell[y*15+x]
}

func (b *Board) PutWord(word string, x, y int, vertical bool) {
	for i := 0; i < len(word); i++ {
		b.Putchar(lc(word[i]), x, y)
		if vertical {
			y++
		} else {
			x++
		}
	}
}

func (b *Board) GetWord(x, y int, vertical bool) string {
	word := ""
	for {
		char := b.Getchar(x, y)
		if char == 0 {
			return word
		}
		word = word + string(char)
		if vertical {
			y++
		} else {
			x++
		}
	}
}

// return the factor to multiply a letter played on x,y by (always 1 if cell already occupied)
func (b *Board) LetterMultiple(x, y int) int {
	return b.Multiple(x, y, &letterMultipleMap)
}

// return the factor to multiply a word played on x,y by (always 1 if cell already occupied)
func (b *Board) WordMultiple(x, y int) int {
	return b.Multiple(x, y, &wordMultipleMap)
}

func (b *Board) Multiple(x, y int, m *map[int]int) int {
	if b.Getchar(x, y) != 0 {
		return 1
	}

	p := y*15 + x
	multiple, exists := (*m)[p]
	if !exists {
		return 1
	}
	return multiple
}

func (b *Board) String() string {
	s := ""
	for y := 0; y < 15; y++ {
		for x := 0; x < 15; x++ {
			if b.Getchar(x, y) == 0 {
				s += "."
			} else {
				s += string(b.Getchar(x, y))
			}
		}
		s += "\n"
	}
	return s
}
