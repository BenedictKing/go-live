package main

import (
	"fmt"
	"strings"
)

func main() {
	correctedSentence := strings.ReplaceAll(sentence, " i ", " I ")
	fmt.Println(correctedSentence)

	aCount := strings.Count(strings.ToLower(goHistory), "a")
	fmt.Println(aCount)

	uppercaseScream := strings.ReplaceAll(text, "scream", "SCREAM")
	fmt.Println(uppercaseScream)
}

// 1. Given the following sentence: Change all of the lowercase letters i, in reference to yourself, to the uppercase I
// Careful not to uppercase the i in the word pair

var sentence = "When i went to the mall i bought a pair of shoes."

// 2. The history of Go is interesting but how many times does the letter a appear? Find a way to count both upper and lower case A's

var goHistory = "Go, not to be confused with the board game Go, was created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. Go was not always known as Go to the public: the language was designed in-house and announced in November 2009. In March of 2012 version 1.0 was released, then later that same year Go 1.1 followed. The domain golang.org was adopted because go.com was already taken. This was somewhat of a practical move at the time, with a short name being very popular around then."

// 3. Given the following sentence: Uppercase the entire word scream each time it appears in the sentence
// I SCREAM, you SCREAM, we all SCREAM for icecream

var text = "I scream, you scream, we all scream for icecream"
