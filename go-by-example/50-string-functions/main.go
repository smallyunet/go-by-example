package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	p("Contains", strings.Contains("test", "es"))
}
