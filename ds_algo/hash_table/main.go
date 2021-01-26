package main

import "fmt"

const ArraySize = 7

// HashTable struct
type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket struct
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}

}

func (b *bucket) search(k string) bool {
	cur := b.head
	for cur != nil {
		if cur.key == k {
			return true
		}
		cur = cur.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	cur := b.head
	for cur.next != nil {
		if cur.next.key == k {
			cur.next = cur.next.next
		}
		cur = cur.next
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		testHashTable.Insert(v)
	}
	testHashTable.Delete("KENNY")

	fmt.Println(testHashTable.Search("KENNY"))
	fmt.Println(testHashTable.Search("BUTTERS"))
}
