///usr/bin/env go run "$0" "$@"; exit "$?"
// the line above is a shebang-like line for go
// chmod +x os_exit_seven.go
// ./os_exit_seven.go
// echo "$?"  # unfortunately it is 1 instead of 7 but at least it is not zero!

package main

import "os"

func main() {
	os.Exit(7)
}
