package main

import (
	"fmt"

	"github.com/atselvan/ankiconnect/ankiconnect"
)

func main() {
	c := ankiconnect.NewClient()
	fmt.Println(c.Deck.GetAll())
	fmt.Println(c.Deck.Create("test"))
	fmt.Println(c.Deck.Delete("test"))
}
