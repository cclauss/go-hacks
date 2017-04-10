///usr/bin/env go run "$0" "$@"; exit "$?"
// the line above is a shebang-like line for go
// chmod +x os_exit_seven.go
// ./os_exit_seven.go
// echo "$?"  # unfortunately it is 1 instead of 7 but at least it is not zero!

package main

import "fmt"
import "net/http"

func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {

func pypi_info(package_name string) (resp *Response, err error) {
    url := 'https://pypi.python.org/pypi/' + package_name
    fmt.Printf("%v\n", url)
    return http.Get(url)

// resp, err := http.Get("http://example.com/")
//if err != nil {
//	// handle error
//}

func main() {
	fmt.Printf("%v (%v)\n", pypi_info('requests'))
}
