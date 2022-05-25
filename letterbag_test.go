package main

import "testing"

func TestBag(t *testing.T) {
    bag := NewLetterBag()

    // check the count of each letter
    count := make(map[byte]int)
    for _, c := range bag.letters {
        count[c]++
    }
    for c, _ := range letterCount {
        if count[c] != letterCount[c] {
            t.Errorf("expected %d of '%c'; got %d", letterCount[c], c, count[c])
        }
    }

    // naive check that they're shuffled
    sameness := 0
    for i := 0; i < len(bag.letters)-1; i++ {
        if bag.letters[i] == bag.letters[i+1] {
            sameness++
        }
    }
    if sameness > 8 {
        t.Errorf("sameness seems too high (%d); please run the test again, and if it keeps failing then fix the shuffle", sameness)
    }
}
