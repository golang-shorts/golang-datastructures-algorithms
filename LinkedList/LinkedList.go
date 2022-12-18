package main

import (
	"encoding/json"
	"fmt"
	"errors"
)

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

func main() {
	var ll LinkedList
	ll.append(newNode("Rajesh"))
	ll.append(newNode("Gowthaman"))
	ll.append(newNode("Subha"))
	ll.append(newNode("Sathish"))

	fmt.Println(ll.get("Lokkesh"))
	fmt.Println(ll.remove("Rajesh"))
	
	fmt.Println(ll.insertAt(newNode("Lokkesh"), 0))
	fmt.Println(ll.get("Lokkesh"))
	fmt.Println(ll.remove("Lokkesh"))
	ll.iterateNodes()
}

func (ll *LinkedList) insertAt(newNode *Node, position int) (isInserted bool, err error) {
	
	if ll.head == nil {
		if position == 0 {
			ll.head = newNode
			ll.tail = ll.head
			isInserted = true
		} else {
			err = errors.New("invalid position")
		}
	} else {
		if position == 0 {
			newNode.next = ll.head
			ll.head = newNode
			isInserted = true
		} else {
			index := 0
			iterator := ll.head
			for ; iterator != nil; iterator = iterator.next {
				index++
				if position == index {
					newNode.next = iterator.next
					iterator.next = newNode
					isInserted = true
					break
				}
				
			}
		}
	}
	return
}

func (ll *LinkedList) get(value string) (node *Node, err error) {
	if ll.head == nil {
		return nil, errors.New("Zero items in list")
	}
	iterator := ll.head

	for ; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator, nil
		}
	}
	return nil, errors.New("Value not found")
}

func (ll *LinkedList) remove(value string) (isRemoved bool, err error) {
	if ll.head == nil {
		return false, errors.New("zero items in list")
	}
	if ll.head.value == value {
		ll.head = ll.head.next
		return true, nil
	}
	iterator := ll.head
	previous := ll.head
	for ; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			previous.next = iterator.next
			return true, nil 			
		}
		previous = iterator
	}
	return false, errors.New("Value not found")
}
func (ll *LinkedList) append(newNode *Node) {
	if ll.head == nil {
		ll.head = newNode
		ll.tail = ll.head
	} else {
		ll.tail.next = newNode
		ll.tail = ll.tail.next
	}
}

func (ll LinkedList) String() {
	value, _ := json.Marshal(ll)
	fmt.Println(string(value))
}

func newNode(value string) *Node {
	var node *Node = new(Node)
	node.value = value
	return node
}

func (ll LinkedList) iterateNodes() {
	iterator := ll.head
	for ; iterator != nil; iterator = iterator.next {
		fmt.Println(iterator.value)
	}
}
func (node Node) String() string {
	return fmt.Sprint("Value is ", node.value)
}
