package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	var hello_message string = "Hello, World!"
	fmt.Println(stringutil.Reverse(hello_message))
}
