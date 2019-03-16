package iter2

// Iterable thing
type Iterable interface {
	Iter() Iterator
}

// LinkedList one
type LinkedList struct {
	node *aNode
}

// NewLinkedList two
func NewLinkedList() *LinkedList {
	node := &aNode{Val: 1}
	node2 := &aNode{Val: 2}
	node2.Add(&aNode{Val: 3})
	node.Add(node2)
	l := &LinkedList{node: node}
	return l
}

// Iter thing
func (linkedList *LinkedList) Iter() Iterator {
	return linkedList.node
}

// Iterator thing
type Iterator interface {
	// Next return true if there is one
	Next() Iterator
}

// Node a type that can be used in iteration
type aNode struct {
	next *aNode
	Val  int
}

// Next returns the next element and an int
// func (node *Node) Next() (*Node, bool) {
// 	return node.next, node.next != nil
// }
func (node *aNode) Next() Iterator {
	return node.next
}

// Add add the next node
func (node *aNode) Add(another *aNode) {
	node.next = another
}
