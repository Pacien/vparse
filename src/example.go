package main

import (
	"encoding/json"
	"fmt"
	"os"
	"vparse"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func vToJson(filename string) string {
	f, err := os.Open(filename)
	check(err)

	parsed := vparse.Parse(f)

	jsonEncoded, err := json.MarshalIndent(parsed, "", "\t")
	check(err)

	return string(jsonEncoded)
}

func main() {
	fmt.Println(vToJson("vcal.ics"))
	fmt.Println(vToJson("vcard.vcs"))
}
