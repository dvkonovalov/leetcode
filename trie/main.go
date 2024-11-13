package trie

import "fmt"

func RunTask() {
	trie := Constructor()
	fmt.Println()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple")) // return True
	fmt.Println(trie.Search("app"))   // return False
	//fmt.Println(trie.StartsWith("app")) // return True
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // return True
}
