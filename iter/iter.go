package iter

// Iterable thing
type Iterable interface {
	Iter() Iterator
}

// Iterator An interface that your iterator should implement
type Iterator interface {
	// Next return true if there is one
	Next() bool
	Value() interface{}
}

type iterFunc func(Iterable)

// Each apply fn to each element of the iterable
func Each(iterable Iterable, fn iterFunc) {
	for iterator := iterable.Iter(); iterator.Next(); {
		fn(iterator)
	}
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

type linkedListIterator struct {
	node *aNode
}

// Iter thing
func (linkedList *LinkedList) Iter() Iterator {
	fakeNode := &aNode{next: linkedList.node, Val: 0}
	return &linkedListIterator{node: fakeNode}
}

// Iterator thing

// Node a type that can be used in iteration
type aNode struct {
	next *aNode
	Val  int
}

// Next returns the next element and an int
// func (node *Node) Next() (*Node, bool) {
// 	return node.next, node.next != nil
// }
func (iterator *linkedListIterator) Next() bool {
	iterator.node = iterator.node.next
	return iterator.node != nil
}

func (iterator *linkedListIterator) Value() interface{} {
	return iterator.node.Val
}

// Add add the next node
func (node *aNode) Add(another *aNode) {
	node.next = another
}
