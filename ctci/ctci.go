package main

import (
	"fmt"
)

func main() {
	uniqueCharacters("uniqe")
	uniqueCharacters("hello")
	reverseString("hello")
}

func uniqueCharacters(word string) bool {
	// we deal with bytes not characters
	// becuase there's no basic char
	// primitive in Go. There is a
	// utf-8 type, but this is simpler

	// also, Go doesn't have built in
	// sets. Our map is going to mape from
	// character-byte to garbage-byte (we don't
	// care at all about the value) where the
	// char exists if it's in the keyset
	charset := make(map[byte]byte)

	for i := 0; i < len(word); i++ {
		_, exists := charset[word[i]]

		if exists {
			fmt.Println(word, "has duplicates")
			return false
		}

		charset[word[i]] = 1
	}
	fmt.Println(word, "has no duplicates")
	return true
}

func reverseString(word string) string {
	// Arrays in Go must be constant size, so we use
	// a slice instead. make() is a special
	// function only for slices, maps, and chans. We use
	// it here to create a slice of type byte with size
	// 'len(word)' and default capacity
	buffer := make([]byte, len(word))

	for i := 0; i < len(word); i++ {
		buffer[len(word) - 1 - i] = word[i]
	}

	fmt.Println("Original word,", word, "reversed:", string(buffer))

	// I believe the conversion is implemented to find the slice's
	// array and copy it into a string (arrays are values in Go, not
	// pointers)
	return string(buffer)
}