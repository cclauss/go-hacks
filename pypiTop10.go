package main

import (
	"fmt"
	"log"

	"github.com/kolo/xmlrpc"
)

func main() {
	client, err := xmlrpc.NewClient("https://pypi.python.org/pypi", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var packages [][]interface{}
	if err := client.Call("top_packages", 1000, &packages); err != nil {
		log.Fatal(err)
	}
	m := make(map[string]int64)
	// var pkgInfos []pkgInfo
	for _, p := range packages {
		m[p[0].(string)] = p[1].(int64)
	}
	fmt.Println(len(packages), len(m), m)
}
