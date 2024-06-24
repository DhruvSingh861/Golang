package main

import "fmt"

type Set struct {
	elements map[string]struct{}
}

func newSet() Set {
	return Set{
		elements: make(map[string]struct{}),
	}
}
func (set Set) add(s string) {
	set.elements[s] = struct{}{}
}
func (set Set) remove(s string) {
	delete(set.elements, s)
}
func (set Set) contains(s string) bool {
	_, a := set.elements[s]
	return a
}

func main() {
	fmt.Println("dshknnjasdjhfgdjhdeukijaxnzmoiuyweqsdgvkjlzxcv")
	set := newSet()
	set.add("dhruv")
	set.add("singh")
	fmt.Println(set)
	set.remove("dhruv")
	fmt.Println(set.contains("dhruv"))
	fmt.Println(set.contains("singh"))
}
