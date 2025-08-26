/*
Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

WordDictionary() Initializes the object.
void addWord(word) Adds word to the data structure, it can be matched later.
bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.
 

Example:

Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True
 

Constraints:

1 <= word.length <= 25
word in addWord consists of lowercase English letters.
word in search consist of '.' or lowercase English letters.
There will be at most 2 dots in word for search queries.
At most 104 calls will be made to addWord and search


Implement above in golang
*/


package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNode
}

// Constructor
func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Add word to the Trie
func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search a word with support for '.' as wildcard
func (this *WordDictionary) Search(word string) bool {
	return searchHelper(word, 0, this.root)
}

// Recursive helper for wildcard support
func searchHelper(word string, index int, node *TrieNode) bool {
	if index == len(word) {
		return node.isEnd
	}

	ch := rune(word[index])
	if ch == '.' {
		for _, child := range node.children {
			if searchHelper(word, index+1, child) {
				return true
			}
		}
		return false
	} else {
		child, exists := node.children[ch]
		if !exists {
			return false
		}
		return searchHelper(word, index+1, child)
	}
}

// Test
func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")

	fmt.Println(wordDictionary.Search("pad")) // false
	fmt.Println(wordDictionary.Search("bad")) // true
	fmt.Println(wordDictionary.Search(".ad")) // true
	fmt.Println(wordDictionary.Search("b..")) // true
}
