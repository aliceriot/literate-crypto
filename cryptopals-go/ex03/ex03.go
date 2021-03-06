/*
---
layout: matasano_exercise
title: "Exercise 3"
key: "ex03"
---

# Exercise 3

[Matasano exercise 3](http://cryptopals.com/sets/1/challenges/3/). This asks
us to break a single byte XOR. Basically, this means that we have a
ciphertext which has been XORed against a single byte. This means we
could exhaustively try every byte until we get it, but we're going to at least
try to do something a little more elegant.
*/

package main

import (
	"encoding/hex"
	"fmt"
	"strings"
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
*/

/*
We'll also need two little helper functions. They each take an XORed byte
array and check that the result has some attribute. The first checks that
the result contains only printable ASCII characters, and the second checks
that the result contains at least one space character (an assumption we make
about the content of the plaintext):
*/

func asciiCheck(bytes []byte) bool {
	for _, c := range bytes {
		if c > 127 || c < 32 {
			return false
		}
	}
	return true
}

func spaceCheck(bytes []byte) bool {
	for _, c := range string(bytes) {
		if c == ' ' {
			return true
		}
	}
	return false
}

/*
We're going to also make one more assumption about the plaintext: we assume
that the most common byte in a decoded plaintext will correspond to one of
the letters `EAOT` or a space, since these are the four most common English
letters. Here's two functions that check for this scenario. The first takes
a string and returns a `map[rune]int` giving the character counts in the string:
*/

func charCount(str string) map[rune]int {
	counts := make(map[rune]int)
	for _, c := range str {
		counts[c]++
	}
	return counts
}

/*
The second function checks that the most common byte in a putative plaintext
is one of the characters `EAOT `. To do this we get the character count of a
lowercased plaintext, find the most common character, and then check to see
if that is one of our key characters:
*/

func aeotCheck(bytes []byte) bool {
	counts := charCount(strings.ToLower(string(bytes)))
	var biggest rune
	count := 0
	for k, v := range counts {
		if v > count {
			biggest = k
			count = v
		}
	}
	for _, c := range "aeto " {
		if biggest == c {
			return true
		}
	}
	return false
}

/*
Great! Together this lets us define a criterion for a valid plaintext:
*/

func validPlaintext(plain []byte) bool {
	return asciiCheck(plain) && spaceCheck(plain) && aeotCheck(plain)
}

/*
## Doing the work

Now that we've written all of the functions we'll need, we can try to
solve the problem!

Basically, we're dealing with a keyspace of all the values a single byte
can take on (0 - 255). This isn't so many, so while we don't want to just
manually check each one, it won't be too much work to check the result of
'decrypting' with each value for certain attributes.

So, we're basically going to iterate through all the possible keys, attempt
to decrypt, and then check `validPlaintext`:
*/

func main() {
	const cipherText = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cipherBytes, _ := hex.DecodeString(cipherText)

	keysAndResults := make(map[byte]string)

	for i := byte(0); i < 255; i++ {
		plain := arrayXOR(cipherBytes, i)
		if validPlaintext(plain) {
			keysAndResults[i] = string(plain)
		}
	}

	for k, plain := range keysAndResults {
		fmt.Printf("key: %d\tplaintext: %s\n", k, plain)
	}
}

/*
Great! After all that we get one key out the end, and a readable plaintext.
*/
