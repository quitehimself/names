// Command generateNames prints a randomly-generated name to stdout, consisting
// of a 3-syllable adverb, a 3-syllable adjective, and a 2-syllable noun.
package main

import (
	"fmt"
	"github.com/quitehimself/names"
)

// TODO: Add word-delimiter and line-delimiter flags and an option
// to output more than one name

func main() {
	fmt.Println(names.Generate(" "))
}
