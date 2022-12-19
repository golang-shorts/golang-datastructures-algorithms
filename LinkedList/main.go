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
	ll.Append(NewNode("g"))
	ll.Append(NewNode("h"))

	fmt.Println(ll.Middle())

}
