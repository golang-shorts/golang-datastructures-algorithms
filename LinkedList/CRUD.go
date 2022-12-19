package main

import (
	"errors"
	"fmt"
)

func (ll *LinkedList) InsertAt(newNode *Node, position int) (isInserted bool, err error) {
	if ll.Head == nil {
		if position == 0 {
			ll.Head = newNode
			ll.Tail = ll.Head
			return true, nil
		} else {
			err = errors.New("invalid position")
		}
	} else {
		if position == 0 {
			newNode.Next = ll.Head
			ll.Head = newNode
			return true, nil
		} else {
			index := 0
			iterator := ll.Head
			for ; iterator != nil; iterator = iterator.Next {
				index++
				if position == index {
					newNode.Next = iterator.Next
					iterator.Next = newNode
					if position == ll.Count {
						ll.Tail = newNode
					}
					return true, nil
					break
				}

			}
		}
	}
	return false, errors.New("unable to insert at a given position")
}

func (ll *LinkedList) Get(value string) (node *Node, err error) {
	if ll.Head == nil {
		return nil, errors.New("Zero items in list")
	}
	iterator := ll.Head

	for ; iterator != nil; iterator = iterator.Next {
		if iterator.Value == value {
			return iterator, nil
		}
	}
	return nil, errors.New("value not found")
}

func (ll *LinkedList) Remove(value string) (isRemoved bool, err error) {
	if ll.Head == nil {
		return false, errors.New("zero items in list")
	}
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		ll.Count--
		return true, nil
	}
	iterator := ll.Head
	previous := ll.Head
	for ; iterator != nil; iterator = iterator.Next {
		if iterator.Value == value {
			previous.Next = iterator.Next
			ll.Count--
			return true, nil
		}
		previous = iterator
	}
	return false, errors.New("value not found")
}

func (ll *LinkedList) Append(newNode *Node) {
	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = ll.Head
	} else {
		ll.Tail.Next = newNode
		ll.Tail = ll.Tail.Next
	}
	ll.Count++
}

func (ll LinkedList) IterateNodes() {
	iterator := ll.Head
	fmt.Println()
	for ; iterator != nil; iterator = iterator.Next {
		fmt.Print(" ", iterator.Value)
	}
	fmt.Println()

}
