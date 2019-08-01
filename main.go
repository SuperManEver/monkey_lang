package main

import (
	"fmt"

	"github.com/SuperManEver/monkey_lang/token"
)

var persons = map[string]string{
	"John": "Doe",
	"Bill": "Atkinson",
}

func main() {
	// str := token.EOF

	fmt.Printf("hello world %s\n", token.LPAREN)
	fmt.Printf("Person: %s\n", persons["John"])
}
