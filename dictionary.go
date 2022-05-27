package main

import (
	"bufio"
	"os"
)

// TODO: do we want to be able to know how popular each word is?

type Dictionary struct {
	words []string
}

func NewDictionary() *Dictionary {
	return &Dictionary{}
}

func (d *Dictionary) AddWord(word string) {
	d.words = append(d.words, word)
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
	for _, w := range d.words {
		if w == word {
			return true
		}
	}
	return false
}
