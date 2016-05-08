/*
---
layout: matasano_exercise
title: "Exercise 1"
---

# Exercise 3

[Matasano exercise 3](http://cryptopals.com/sets/1/challenges/3/). This asks
us to break a single byte XOR. Basically, this means that we have a
ciphertext which has been XORed against a single byte. This means we
can exhaustively try every byte until we get it.
*/

package main

import (
	"encoding/hex"
	"fmt"
)

/*
First off, something we're almost certainly going to need is a function
that takes a byte slice and a byte, and XORs the contents of that byte slice
with that byte.

We'll call it `arrayXOR`:
*/

func arrayXOR(in []byte, n byte) []byte {
	out := make([]byte, len(in))
	for i, v := range in {
		out[i] = v ^ n
	}
	return out
}

/*
Then if we want to XOR a particular byte, say, 42, with a byte array we can
just do `xorResult := arrayXOR(in, byte(42))`. Nice!

Next we'll need a function that takes a string and returns a count
of the occurences of each letter.
We can use a byte -> int map to keep track.
*/

func charCount(bytes []byte, counts map[byte]int) {
	for _, c := range bytes {
		counts[c]++
	}
}

/*
We'll also need two little helper functions. They each take an XORed byte
array and check that the result has some attribute. The first checks that
the result contains only printable ASCII characters, and the second checks
that the result contains at least one space character (as assumption we make
about the content of the plaintext).
*/

func asciiCheck(bytes []byte) bool {
	for _, c := range bytes {
		if int(c) > 127 {
			return false
		}
	}
	return true
}

func spaceCheck(bytes []byte) bool {
	found := false
	for _, c := range string(bytes) {
		if c == ' ' {
			found = true
		}
	}
	return found
}

/*
## Doing the work

now that we've written all of the functions we'll need, we can go ahead
and solve the problem!

First, declare the ciphertext
*/

func main() {
	const cipherText = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cipherBytes, _ := hex.DecodeString(cipherText)

	possibleKeys := make([]byte, 0)
	for i := byte(0); i < 255; i++ {
		plain := arrayXOR(cipherBytes, i)
		if asciiCheck(plain) && spaceCheck(plain) {
			possibleKeys = append(possibleKeys, byte(i))
		}
	}

	fmt.Println(possibleKeys)

	// 	counts := make(map[byte]int)
	// 	charCount(cipherBytes, counts)

	// 	out := make([]byte, len(cipherBytes))
	// 	arrayXOR(cipherBytes, out, byte(42))
	// 	fmt.Println(out)

}
