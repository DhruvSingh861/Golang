package main

import (
	"container/list"
	"fmt"
	"sync"
)

func main() {

	// Simple Map implementation
	map1 := make(map[string]string)
	map1["s1"] = "Dhruv"
	map1["s2"] = "singh"
	fmt.Println(map1)

	// Synchronized Map
	// It is safe for multiple goroutines to call a Map's methods concurrently.
	var synchronizedMap sync.Map
	synchronizedMap.Store("name", "dhruv")
	synchronizedMap.Store("surname", "Singh")
	synchronizedMap.Store("name", "Dhruv")
	synchronizedMap.Store("Id", 10)
	temp1, _ := synchronizedMap.Load("name")
	fmt.Println("synchronized Map value for Key name: ", temp1)

	// The Go language doesn’t have any sets by default like we have
	// in other languages but there are ways of implementing a golang set.
	// Set impl. using Map
	// creating empty structs require zero memory.
	set := make(map[string]struct{})
	set["S1"] = struct{}{}
	set["S2"] = struct{}{}
	set["S3"] = struct{}{}
	set["S4"] = struct{}{}
	set["S4"] = struct{}{}

	delete(set, "S1")
	_, a := set["S2"]
	fmt.Printf("contains S2 : %t\n", a)
	_, b := set["S8"]
	fmt.Printf("contains S8 : %t\n", b)

	//List implementations
	l := list.New()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.PushBack(40)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i)
	}

}
