package datastructures

type Node struct {
	Value any
	next  *Node
}

type LinkedList struct {
	head, tail, trav *Node
}

func New() (linkedList LinkedList) {
	return
}

func (linkedList *LinkedList) PushBack(data any) {
	if linkedList.head != nil {
		linkedList.tail.next = &Node{Value: data}
		linkedList.tail = linkedList.tail.next
	} else {
		linkedList.head = &Node{Value: data}
		linkedList.trav = linkedList.head
		linkedList.tail = linkedList.head
	}
}

func (linkedList *LinkedList) Front() (front *Node) {
	front = linkedList.head
	return
}

// RemoveFirst O(1)
func (linkedList *LinkedList) RemoveFirst() {
	linkedList.head = linkedList.head.next
	linkedList.trav = linkedList.head
}

// RemoveLast O(n)
func (linkedList *LinkedList) RemoveLast() {
	var temp *Node
	println()
	for e := linkedList.Front(); e != nil; e = linkedList.Next() {
		if e.Value == linkedList.tail.Value {
			linkedList.trav = linkedList.head
			break
		}
		temp = e
	}
	temp.next = nil
	linkedList.tail = temp
}

// Next O(1)
func (linkedList *LinkedList) Next() (node *Node) {
	linkedList.trav = linkedList.trav.next
	if linkedList.trav != nil {
		node = linkedList.trav
	} else {
		linkedList.trav = linkedList.head
	}
	return
}
