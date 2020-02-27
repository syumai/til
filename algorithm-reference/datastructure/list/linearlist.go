package list

import "fmt"

var (
	ErrLinearListNodeOutOfBounds = fmt.Errorf("index out of bounds")
)

type LinearListNode struct {
	Value int
	Next  *LinearListNode
}

func (node *LinearListNode) Add(i int, value int) error {
	if i < 0 || i > node.Length()-1 {
		return ErrLinearListNodeOutOfBounds
	}
	n := LinearListNode{
		Value: value,
	}

	// insert to top
	if i == 0 {
		n.Next = node.Next

		tmpVal := node.Value
		node.Value = n.Value
		n.Value = tmpVal

		node.Next = &n
		return nil
	}

	// insert to i
	a, err := node.Get(i - 1)
	if err != nil {
		return err
	}
	b := a.Next
	a.Next = &n
	n.Next = b
	return nil
}

func (node *LinearListNode) Get(i int) (*LinearListNode, error) {
	if i < 0 || i > node.Length()-1 {
		return nil, ErrLinearListNodeOutOfBounds
	}
	n := node
	j := 0
	for {
		if i == j {
			break
		}
		n = n.Next
		j++
	}
	return n, nil
}

func (node *LinearListNode) Length() int {
	l := 1
	n := node
	for {
		n = n.Next
		if n == nil {
			break
		}
		l++
	}
	return l
}

func (node *LinearListNode) ForEach(f func(*LinearListNode)) {
	n := node
	for {
		f(n)
		n = n.Next
		if n == nil {
			break
		}
	}
}

func (node *LinearListNode) ToSlice() []*LinearListNode {
	var s []*LinearListNode
	node.ForEach(func(n *LinearListNode) {
		s = append(s, n)
	})
	return s
}
