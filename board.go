package main

type Board struct {
    cell [225]byte
}

var wordMultipleMap = map[int]int {
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
var letterMultipleMap = map[int]int {
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
var letterScore = map[byte]int {
    'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1,
    'f': 4, 'g': 2, 'h': 4, 'i': 1, 'j': 8,
    'k': 5, 'l': 1, 'm': 3, 'n': 1, 'o': 1,
    'p': 3, 'q': 10,'r': 1, 's': 1, 't': 1,
    'u': 1, 'v': 4, 'w': 4, 'x': 8, 'y': 4,
    'z': 10,
}

func NewBoard() Board {
    return Board{}
}

// return the score for playing the given word at the given position, with no legality check;
// words should be all lowercase, except for blank tiles which should be uppercase
func (b *Board) Score(word string, x,y int, vertical bool) int {
    score := 0

    wordFactor := 1

    for i := 0; i < len(word); i++ {
        if b.Getchar(x,y) == 0 {
            score += b.LetterMultiple(x, y) * letterScore[word[i]]
            wordFactor *= b.WordMultiple(x, y)
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
func (b *Board) Play(word string, x,y int, vertical bool) int {
    score := b.Score(word, x, y, vertical)

    // lowercase for ascii
    lc := func(c byte) byte {
        if c >= 'A' && c <= 'Z' {
            return c + 'a' - 'A'
        } else {
            return c
        }
    }

    for i := 0; i < len(word); i++ {
        b.Putchar(lc(word[i]), x, y)
        if vertical {
            y++
        } else {
            x++
        }
    }

    return score
}

func (b *Board) Putchar(char byte, x,y int) {
    if x < 0 || x >= 15 || y < 0 || y >= 15 {
        return
    }
    b.cell[y*15+x] = char
}

func (b *Board) Getchar(x,y int) byte {
    if x < 0 || x >= 15 || y < 0 || y >= 15 {
        return 0
    }
    return b.cell[y*15+x]
}

// return the factor to multiply a letter played on x,y by (always 1 if cell already occupied)
func (b *Board) LetterMultiple(x,y int) int {
    return b.Multiple(x,y, &letterMultipleMap)
}

// return the factor to multiply a word played on x,y by (always 1 if cell already occupied)
func (b *Board) WordMultiple(x,y int) int {
    return b.Multiple(x,y, &wordMultipleMap)
}

func (b *Board) Multiple(x,y int, m *map[int]int) int {
    if b.Getchar(x,y) != 0 {
        return 1
    }

    p := y*15+x
    multiple, exists := (*m)[p]
    if !exists {
        return 1
    }
    return multiple
}
