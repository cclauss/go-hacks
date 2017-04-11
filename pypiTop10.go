///usr/bin/env go run "$0" "$@"; exit "$?"
// the line above is a shebang-like line for go
// chmod +x os_exit_seven.go
// ./os_exit_seven.go
// echo "$?"  # unfortunately it is 1 instead of 7 but at least it is not zero!

package main

import (
	"fmt"
	"log"
	"reflect"

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
	fmt.Println(len(packages), packages) // 10 (correct)
	for _, p := range packages {
		fmt.Println(reflect.TypeOf(p), p)  // []interface {} [six 110953835]
		// How do I get just the p.pkgName ("six" as a string)?
		// How do I get just the p.downloads (110953835 as an int)?
	}
}
