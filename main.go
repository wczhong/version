package main

import (
	"flag"
	"fmt"

	"example/mysql/version"
)

var ver = flag.Bool("version", false, "version")

func main() {
	// flag.Parse()
	fmt.Println(version.String())
}
