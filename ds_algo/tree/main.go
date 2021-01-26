package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Tree struct {
	root *Node
}

func InitTree() *Tree {
	result := &Tree{
		root: &Node{},
	}
	return result
}

func (t *Tree) Insert(w string) {
	wordLen := len(w)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a' // because we want a to become 0
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Tree) Search(w string) bool {
	wordLen := len(w)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a' // because we want a to become 0
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd {
		return true
	}
	return false
}

func main() {
	res := InitTree()

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		res.Insert(v)
	}
	// fmt.Println(res.root)
	// res.Insert("aragorn")
	fmt.Println(res.Search("aragosdan"))
}
