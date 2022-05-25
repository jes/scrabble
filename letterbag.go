package main

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
    for c := 'a'; c <= 'z'; c++ {
        for i := 0; i < letterCount[c]; i++ {
            l.letters = append(l.letters, c)
        }
    }
    // shuffle the bag
    TODO: shuffle
	return &l
}
