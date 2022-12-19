package main

func (ll *LinkedList) Reverse() (isDone bool, err error) {

	var prev *Node
	var current *Node = ll.Head
	var iteration int = -1
	for current != nil {
		iteration++
		if iteration == 0 {
			ll.Tail = current
		}
		var next *Node = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	ll.Head = prev
	return false, nil
}
