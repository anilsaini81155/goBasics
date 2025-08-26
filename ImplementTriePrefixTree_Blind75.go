/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.

Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True


Constraints:

1 <= word.length, prefix.length <= 2000
word and prefix consist only of lowercase English letters.
At most 3 * 104 calls in total will be made to insert, search, and startsWith.

Implement above in golang

Base Code:

type Trie struct {

}


func Constructor() Trie {

}


func (this *Trie) Insert(word string)  {

}


func (this *Trie) Search(word string) bool {

}


func (this *Trie) StartsWith(prefix string) bool {

}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
*/

package main

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, ch := range word {
		if _, found := node.children[ch]; !found {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	// fmt.Println("node:", node)
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, ch := range word {
		if _, found := node.children[ch]; !found {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		if _, found := node.children[ch]; !found {
			return false
		}
		node = node.children[ch]
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	println(trie.Search("apple"))   // true
	println(trie.Search("app"))     // false
	println(trie.StartsWith("app")) // true
	trie.Insert("app")
	println(trie.Search("app")) // true
}
