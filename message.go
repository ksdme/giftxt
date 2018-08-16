package giftxt

import "strings"

// Message is a struct indicating the
// structure of an incoming message.
type Message struct {
	Text    string
	Words   []string
	Longest string
}

// GetWordsFromString returns a slice containing
// the words of a given string.
func GetWordsFromString(str string) []string {
	return strings.Split(str, " ")
}

// GetLongestWord returns the longest word
// from a given slice of words.
func GetLongestWord(words []string) string {
	var longest string
	length := -1

	for _, word := range words {
		templen := len([]rune(word))
		if templen > length {
			length = templen
			longest = word
		}
	}

	return longest
}

// ProcessText processes text before making
// a message out of it.
func ProcessText(str string) string {
	return strings.ToUpper(str)
}

// NewMessage returns pointer to new
// initialized Message struct.
func NewMessage(str string) *Message {
	str = ProcessText(str)

	words := GetWordsFromString(str)
	longest := GetLongestWord(words)

	return &Message{
		Text:    str,
		Words:   words,
		Longest: longest,
	}
}
