///usr/bin/env go run "$0" "$@"; exit "$?"
// the line above is a shebang-like line for go
// chmod +x os_exit_seven.go
// ./os_exit_seven.go
// echo "$?"  # unfortunately it is 1 instead of 7 but at least it is not zero!

package main

import (
	"log"

	"github.com/kolo/xmlrpc"
)

// https://warehouse.pypa.io/api-reference/xml-rpc/

func main() {
	client, _ := xmlrpc.NewClient("https://pypi.python.org/pypi", nil)
	defer client.Close()
	var packages []interface{}
	if err := client.Call("top_packages", 10, &packages); err != nil {
		log.Fatal(err)
	}
	println(len(packages)) // 10 (correct)
	type pkgInfo struct {
		Downloads int    `xmlrpc:"downloads"`
		Filename  string `xmlrpc:"filename"`
		URL       string `xmlrps:"url"`
	}
	for _, p := range packages {
		pkg := p.(pkgInfo)  // panic: interface conversion: interface {} is []interface {}, not main.pkgInfo
		println(pkg.Filename)
	}
}
