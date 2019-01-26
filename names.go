// Package names generates names of the form "3-syllable adverb, 3-syllable
// adjective, 2-syllable noun," with the caller's choice of delimiter.
package names // import "github.com/quitehimself/names"
import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// Generate returns a random name consisting of a 3-syllable adverb, a
// 3-syllable adjective, and a 2-syllable noun, separated by the provided
// delimiter.
func Generate(delimiter string) string {
	var seed int64
	binary.Read(crand.Reader, binary.LittleEndian, &seed)
	rand.Seed(seed)
	return adverbs[rand.Intn(len(adverbs))] + delimiter +
		adjectives[rand.Intn(len(adjectives))] + delimiter +
		nouns[rand.Intn(len(nouns))]
}
