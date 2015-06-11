///usr/bin/env go run "$0" "$@"; exit "$?"

package main

import "fmt"

func main() {
	white_pieces := [6]rune{}
	black_pieces := [6]rune{}
	for i := range white_pieces {
		white_pieces[i] = rune('\u2654' + i)
		black_pieces[i] = rune('\u265A' + i)
	}
	fmt.Printf("%c\n%c\n", white_pieces, black_pieces)
}

// [♔ ♕ ♖ ♗ ♘ ♙]
// [♚ ♛ ♜ ♝ ♞ ♟]
