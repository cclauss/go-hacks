package main

/*
+------------------------------------------------------------------------+
|     A  B  C  D  E  F  G  H  I  J  ||  A  B  C  D  E  F  G  H  I  J     |
|  1                                 1                                 1 |
|  2        A  A  A  A  A            2                                 2 |
|  3                                 3                                 3 |
|  4           B  B  B  B            4                                 4 |
|  5                                 5                                 5 |
|  6  C  C  C                        6                                 6 |
|  7                                 7                                 7 |
|  8                       S  S  S   8                                 8 |
|  9                                 9                                 9 |
| 10              D  D              10                                10 |
|     A  B  C  D  E  F  G  H  I  J  ||  A  B  C  D  E  F  G  H  I  J     |
+------------------------------------------------------------------------+
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const a = "  AAAAA   "
const b = "   BBBB   "
const c = "CCC       "
const s = "       SSS"
const d = "    DD    "

const blanks = "          "
const hitAndMiss = "!."
const letters = "ABCDEFGHIJ"

var ships = func() map[string]int {
	ships := make(map[string]int)
	ships["Aircraft Carrier"] = 5
	ships["Battleship"] = 4
	ships["Cruiser"] = 3
	ships["Submarine"] = 3
	ships["Destroyer"] = 2
	return ships
}()

type point struct {
	X, Y int
}

type player struct {
	name  string
	board []string
}

func letterNumberToPoint(s string) (point, error) {
	s = strings.ToUpper(s)
	x := strings.Index(letters, s[:1])
	y := strings.Index("12345678910", s[1:])
	if x == -1 || y == -1 || y == 10 {
		if s[:1] == "Q" {
			panic("User quit.")
		}
		return point{x, y}, errors.New("invalid: Try 'A1' or 'J10'")
	}
	return point{x, y}, nil
}

func borderRow() string {
	return strings.Join([]string{"+", "+"}, strings.Repeat("-", 72))
}

func charsInStr(str string) []string {
	// return [c for c in str]
	letters := []string{}
	for _, c := range str {
		letters = append(letters, string(c))
	}
	return letters
}

func clokeInStr(str string) []string {
	// return [c for c in str]
	letters := []string{}
	for _, c := range str {
		if strings.ContainsRune(hitAndMiss, c) == false {
			c = ' ' // cloke the battleships!
		}
		letters = append(letters, string(c))
	}
	return letters
}

func letterRow() string {
	letters := strings.Join(charsInStr(letters), "  ")
	return fmt.Sprintf("|     %s  ||  %s     |", letters, letters)
}

func formatRow(i int) string {
	return fmt.Sprintf("| %2d  %s  %#2[1]d  %s  %#2[1]d |", i, "%s")
}

func board(homeTeam, awayTeam player) string {
	rows := []string{borderRow(), letterRow()}
	for i := 0; i < 10; i++ {
		h := strings.Join(charsInStr(homeTeam.board[i]), "  ")
		a := strings.Join(clokeInStr(awayTeam.board[i]), "  ")
		rows = append(rows, fmt.Sprintf(formatRow(i+1), h, a))
	}
	return strings.Join(append(append(rows, rows[1]), rows[0]), "\n")
}

func makeHumanPlayer() player {
	return player{"human", []string{a, blanks, b, blanks, c, blanks, s,
		blanks, d, blanks}}
}
func makeCompuPlayer() player {
	return player{"computer", []string{blanks, a, blanks, b, blanks, c,
		blanks, s, blanks, d}}
}

func askWhichSquare(maxSquare string) string {
	fmt.Printf("maxSquare: %s\n", maxSquare)
	reader := bufio.NewReader(os.Stdin)
	// for {
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%T|%s|\n", text, text)
	fmt.Println(letterNumberToPoint(text))
	return text
	// fmt.Println(board(compuPlayer, humanPlayer))
}

func humanTurn() {}
func compuTurn() {}

func main() {
	humanPlayer := makeHumanPlayer()
	compuPlayer := makeCompuPlayer()
	gameOn := true
	for gameOn == true {
		fmt.Println(board(compuPlayer, humanPlayer))
		sq := askWhichSquare("J10")
		if sq[:1] == "Q" {
			gameOn = false
		} else {

			humanTurn()
			compuTurn()
		}
	}
	fmt.Println(board(compuPlayer, humanPlayer))
}
