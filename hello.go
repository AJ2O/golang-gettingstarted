package main

import (
	"fmt"

	"github.com/AJ2O/golang-gettingstarted/morestrings"
	"github.com/google/go-cmp/cmp"
	"rsc.io/quote"
)

func playQuote() {
	fmt.Println(quote.Opt())
}

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	playQuote()
}
