///usr/bin/env go run ${0} ${@}; exit ${?}

package main

import "fmt"

func makeChessPieces() map[string]string {
	var names []string
	for _, color := range "WB" {
		for _, name := range "KQRBNP" {
			names = append(names, string(color)+string(name))
		}
	}
	pieces := make(map[string]string)
	for i, name := range names {
		pieces[name] = string('\u2654' + i)
	}
	return pieces
}

func main() {
	fmt.Printf("%v\n", makeChessPieces())
}

// map[WK:♔ WN:♘ WP:♙ BQ:♛ BR:♜ BN:♞ WQ:♕ WR:♖ WB:♗ BK:♚ BB:♝ BP:♟]
