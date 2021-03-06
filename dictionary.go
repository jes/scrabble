package main

import (
	"bufio"
	"os"
	"strings"
)

// TODO: do we want to be able to know how popular each word is, so that the AI
// can have access to a dictionary of the N most popular words rather than all
// words, to simulate human play?

type Dictionary struct {
	words []string
}

func NewDictionary() *Dictionary {
	return &Dictionary{}
}

// insert 'word' into dictionary in lowercase
func (d *Dictionary) AddWord(word string) {
	d.words = append(d.words, strings.ToLower(word))
}

func (d *Dictionary) AddFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d.AddWord(scanner.Text())
	}
	return nil
}

// XXX: O(n), and likely to be a substantial bottleneck
func (d *Dictionary) HasWord(word string) bool {
	word = strings.ToLower(word)
	for _, w := range d.words {
		if w == word {
			return true
		}
	}
	return false
}
