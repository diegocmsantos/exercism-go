package proverb

import (
	"fmt"
)

var repeatedVerse = "For want of a %s the %s was lost."
var finalVerse = "And all for the want of a %s."

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var ret []string

	if len(rhyme) == 0 {
		return []string{}
	}
	if len(rhyme) == 1 {
		return []string{fmt.Sprintf(finalVerse, rhyme[0])}
	}

	for i := 0; i < len(rhyme)-1; i++ {
		word := rhyme[i]
		nextWord := rhyme[i+1]
		ret = append(ret, fmt.Sprintf(repeatedVerse, word, nextWord))
	}
	ret = append(ret, fmt.Sprintf(finalVerse, rhyme[0]))
	return ret
}
