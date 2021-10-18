package main

import (
    "flag"
	"fmt"
	"github.com/valyala/gozstd"
	"io/ioutil"
)

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadFile("hello_world.zst")
	if err != nil {
		panic(err)
	}

	content, err := gozstd.Decompress(nil, bytes)
	if err != nil {
		panic(err)
	}

	if string(content) != "hello world" {
		panic(content)
	}

	fmt.Printf("Got string from zstd file: %s\n", string(content))
}
