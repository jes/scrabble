package main

// lowercase for ascii
func lc(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 'a' - 'A'
	} else {
		return c
	}
}
