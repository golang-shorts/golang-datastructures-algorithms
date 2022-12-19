package main

import "fmt"

func main() {
	var ll LinkedList
	ll.Append(NewNode("a"))
	ll.Append(NewNode("b"))
	ll.Append(NewNode("c"))
	ll.Append(NewNode("d"))
	ll.Append(NewNode("e"))
	ll.Append(NewNode("f"))
	ll.InsertAt(NewNode("g"), 6)

	fmt.Println(ll.Tail)
	ll.InsertAt(NewNode("h"), 7)

	ll.Reverse()
	ll.IterateNodes()
	ll.Reverse()
	ll.IterateNodes()
	ll.Reverse()
	ll.IterateNodes()
	fmt.Println(ll.Tail)

}
