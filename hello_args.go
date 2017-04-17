///usr/bin/env go run ${0} ${@} ; exit ${?}
// the line above is a shebang-like line for go
// chmod +x hello_args.go
// ./hello_args.go
// echo ${?}  # unfortunately it is 1 instead of 7 but at least it is not zero!

package main

import(
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
    fmt.Printf("Hello, world.\n")
    os.Exit(7)
}
