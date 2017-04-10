package main

import (
	"log"

	"github.com/kolo/xmlrpc"
)

// https://warehouse.pypa.io/api-reference/xml-rpc/

func main() {
	client, _ := xmlrpc.NewClient("https://pypi.python.org/pypi", nil)
	defer client.Close()
	var result [10]interface{}
	err := client.Call("top_packages", 10, &result)
	if err != nil {
		log.Fatal(err)
	}
	// print(result[0].(string))
	for i, r := range result {
		println(i, r)
	}
}
