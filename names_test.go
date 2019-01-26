package names

import (
	"strings"
	"testing"
)

var testChars = []rune{
	// Characters to try as delimiters:
	// plus, minus, underscore, dot, null, and every character from the
	// wikipedia page on whitespace characters
	0x0,
	0x9,
	0xa,
	0xb,
	0xc,
	0xd,
	0x20,
	0x2b,
	0x2d,
	0x2e,
	0x5f,
	0x85,
	0xa0,
	0xb7,
	0x1680,
	0x180e,
	0x2000,
	0x2001,
	0x2002,
	0x2003,
	0x2004,
	0x2005,
	0x2006,
	0x2007,
	0x2008,
	0x2009,
	0x200a,
	0x200b,
	0x200c,
	0x200d,
	0x2028,
	0x2029,
	0x202f,
	0x205f,
	0x2060,
	0x237d,
	0x2420,
	0x2422,
	0x2423,
	0x3000,
	0xfeff,
}

func validateOutput(output string, delimiter rune) bool {
	// Note: this only works because none of the test
	// delimiters is from 'a'-'z', the set of characters
	// used in the word list entries
	words := strings.Split(output, string(delimiter))
	if len(words) != 3 {
		return false
	}
	for i, wordlist := range [3][]string{adverbs, adjectives, nouns} {
		found := false
		for _, word := range wordlist {
			if word == words[i] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestNameGeneration(t *testing.T) {
	for _, delimiter := range testChars {
		output := Generate(string(delimiter))
		if !validateOutput(output, delimiter) {
			t.Fatal("output does not match wordlists.")
		}
	}
}

func TestRandomness(t *testing.T) {
	// With the current number of entries in the lists
	// (200 adverbs, 600 adjectives, and 1200 nouns),
	// the chances of getting the same result twice in
	// a row is 1 in 144,000,000, which is unlikely, but
	// maybe not completely comfortable.  Testing for two
	// duplicates means that, if every person on the planet
	// runs this test once an hour, we would expect to see
	// the test fail roughly once every 300 years.  Those are
	// odds I can live with.
	if (Generate("") == Generate("")) && (Generate("") == Generate("")) {
		t.Fatal("output appears to be nonrandom.")
	}
}
