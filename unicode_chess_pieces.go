///usr/bin/env go run ${0} ${@}; exit ${?}

package main

import "fmt"

func makeChessPieces() map[string]string {
	pieces := make(map[string]string)
	for c, color := range "WB" {
		for n, name := range "KQRBNP" {
			pieces[string(color)+string(name)] = string('\u2654' + c*6 + n)
		}
	}
	return pieces
}

func main() {
	fmt.Printf("%v\n", makeChessPieces())
}

// map[BB:♝, BK:♚, BN:♞, BP:♟, BQ:♛, BR:♜,
//     WB:♗, WK:♔, WN:♘, WP:♙, WQ:♕, WR:♖]
