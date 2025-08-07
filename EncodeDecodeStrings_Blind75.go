/*
Design an algorithm to encode a list of strings to a string. The encoded string is sent over the network and then decoded back to the original list of strings at the receiver's end. Both encode and decode must be stateless and handle any ASCII character.

You are not permitted to use built-in serialization methods such as eval.

Example:

Input:  ["Hello", "World"]
Output: ["Hello", "World"]

Constraints:

1 ≤ strs.length ≤ 200

0 ≤ strs[i].length ≤ 200

Each strs[i] can contain any of the 256 ASCII characters. 

Implement above in golang
*/


package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Encodes a slice of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var buf bytes.Buffer
	for _, s := range strs {
		// Write fixed 4‑digit length, zero‑padded
		buf.WriteString(fmt.Sprintf("%04d", len(s)))
		buf.WriteString(s)
	}
	return buf.String()
}

// Decodes a single encoded string back into the original slice.
// Adds error handling for invalid/corrupted input strings.
func (codec *Codec) Decode(s string) []string {
	var result []string
	i := 0
	n := len(s)
	for i < n {
		if i+4 > n {
			fmt.Println("Warning: Corrupted input — length header is incomplete.")
			break
		}
		length, err := strconv.Atoi(s[i : i+4])
		if err != nil {
			fmt.Println("Warning: Invalid length header.")
			break
		}
		i += 4
		if i+length > n {
			fmt.Println("Warning: Corrupted input — not enough characters for string of length", length)
			break
		}
		result = append(result, s[i:i+length])
		i += length
	}
	return result
}

// Test different input cases
func testCodec(codec Codec, input []string) {
	fmt.Println("Original:", input)
	enc := codec.Encode(input)
	fmt.Println("Encoded :", enc)
	dec := codec.Decode(enc)
	fmt.Println("Decoded :", dec)
	fmt.Println("--------------------------------------------------")
}

func main() {
	codec := Constructor()

	testCases := [][]string{
		{"Hello", "World"},
		{"", "", ""},
		{"Hello\nWorld", "😀", "!@#$%^&*()", "line1\nline2", "tab\tseparated"},
		{"你好", "こんにちは", "안녕하세요", "😀😁😂🤣😃"},
		{"a", "ab", "abc", "abcd", "abcde", "abcdef"},
		{string(bytes.Repeat([]byte("x"), 1000))},
		func() []string {
			// Generate 100 short strings: "0" to "99"
			input := make([]string, 100)
			for i := range input {
				input[i] = strconv.Itoa(i)
			}
			return input
		}(),
		{"1234", "0000", "9876543210"},
		{"0004", "1234test", "0000data"},
	}

	for _, test := range testCases {
		testCodec(codec, test)
	}

	// Optionally test a corrupted encoded string
	fmt.Println("Testing corrupted encoded string:")
	corrupted := "0005Hello0003ABC000" // last string is incomplete
	fmt.Println("Encoded (corrupted):", corrupted)
	decoded := codec.Decode(corrupted)
	fmt.Println("Decoded:", decoded)
}
