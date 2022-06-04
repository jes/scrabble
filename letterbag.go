package main

import "math/rand"

type LetterBag struct {
	letters []byte
}

var letterCount = map[byte]int{
	'a': 9, 'b': 2, 'c': 2, 'd': 4, 'e': 12,
	'f': 2, 'g': 3, 'h': 2, 'i': 9, 'j': 1,
	'k': 1, 'l': 4, 'm': 2, 'n': 6, 'o': 8,
	'p': 2, 'q': 1, 'r': 6, 's': 4, 't': 6,
	'u': 4, 'v': 2, 'w': 2, 'x': 1, 'y': 2,
	'z': 1, ' ': 2,
}

func NewLetterBag() *LetterBag {
	l := LetterBag{}
	// put all the letters in the bag
	for c, num := range letterCount {
		for i := 0; i < num; i++ {
			l.AppendLetter(c)
		}
	}
	// shuffle the bag
	l.Shuffle()
	return &l
}

func (l *LetterBag) Empty() {
	l.letters = make([]byte, 0)
}

//  https://golang.cafe/blog/how-to-shuffle-a-slice-in-go.html
func (l *LetterBag) Shuffle() {
	rand.Shuffle(len(l.letters), func(i, j int) {
		l.letters[i], l.letters[j] = l.letters[j], l.letters[i]
	})
}

func (l *LetterBag) AppendLetter(c byte) {
	l.letters = append(l.letters, c)
	l.Shuffle()
}

func (l *LetterBag) BagSize() int {
	return len(l.letters)
}

func (l *LetterBag) BagEmpty() bool {
	return l.BagSize() == 0
}

// return the next letter, or 0 if the bag is empty
func (l *LetterBag) NextLetter() byte {
	if l.BagEmpty() {
		return 0
	}

	letter := l.letters[0]
	l.letters = l.letters[1:]
	return letter
}

func (l *LetterBag) String() string {
	s := ""
	for i := range l.letters {
		s += string(l.letters[i])
	}
	return s
}
