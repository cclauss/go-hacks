///usr/bin/env go run "$0" "$@"; exit "$?"
// the line above is a shebang-like line for go
// chmod +x os_exit_seven.go
// ./os_exit_seven.go
// echo "$?"  # unfortunately it is 1 instead of 7 but at least it is not zero!

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {

func pypiHTML(packageName string) (string, error) {
	url := "https://pypi.python.org/pypi/" + packageName
	fmt.Printf("%v\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes), err
}

func pypiInfo(packageName string) (map[string]interface{}, error) {
	url := "https://pypi.python.org/pypi/" + packageName + "/json"
	// fmt.Printf("%v\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jsonData interface{}
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		log.Fatal(err)
	}
	jsonData = jsonData.(map[string]interface{})["info"]
	return jsonData.(map[string]interface{}), err
}

func pypiRecord(ch chan string, packageName string) {
	info, err := pypiInfo(packageName)
	if err != nil {
		log.Fatal(err)
	}
	py2Only := false
	py3Support := false
	for _, classifier := range info["classifiers"].([]interface{}) {
		s := classifier.(string)
		// println(i, s)
		if s == "Programming Language :: Python :: 2 :: Only" {
			py2Only = true
		} else if strings.HasPrefix(s, "Programming Language :: Python :: 3") {
			py3Support = true
		}
	}
	r := []string{
		strconv.FormatBool(py2Only),
		strconv.FormatBool(py3Support),
		info["version"].(string),
		info["package_url"].(string),
	}
	ch <- strings.Join(r, ",")
	// return strings.Join(r, ","), err
}

func main() {
	ch := make(chan string)
	packageNames := []string{"requests", "aiohttp", "pip", "ansible", "supervisor", "newrelic_plugin_agent"}
	for _, packageName := range packageNames {
		//		s := strings.Join(pypiRecord(packageName), ",")
		go pypiRecord(ch, packageName)
		// if err != nil {
		//	log.Fatal(err)
		//}
		//println(s)
	}

	for i := range packageNames {
		println(i, <-ch)
	}
}
