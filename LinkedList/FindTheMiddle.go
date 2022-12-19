package main

func (ll *LinkedList) Middle() *Node {
	var runner, fastRunner *Node = ll.Head, ll.Head
	if runner == nil {
		return runner
	}
	for runner.Next != nil && fastRunner != nil && fastRunner.Next != nil {
		runner = runner.Next
		fastRunner = fastRunner.Next.Next
	}
	return runner
}
