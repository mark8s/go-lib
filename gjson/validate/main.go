package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func main() {
	const json = `{"name":"dj",age:18}`
	fmt.Println(gjson.Get(json, "name"))

	if !gjson.Valid(json) {
		fmt.Println("error")
	} else {
		fmt.Println("ok")
	}

	name := "profiles/default.yaml"
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	fmt.Println(cannonicalName)
}
