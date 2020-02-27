package list

import "fmt"

var (
	ErrLinearListOutOfBounds = fmt.Errorf("index out of bounds of linear list")
	ErrLinearListHasNoEntry  = fmt.Errorf("linear list has no entry")
)

type (
	LinearList struct {
		FirstNode *LinearListNode
	}

	LinearListNode struct {
		Value int
		Next  *LinearListNode
	}
)

func (list *LinearList) Add(i int, value int) error {
	if i < 0 || i > list.Length() {
		return ErrLinearListOutOfBounds
	}
	n := LinearListNode{
		Value: value,
	}

	// insert to blank list
	if list.FirstNode == nil {
		list.FirstNode = &n
		return nil
	}

	// insert to top
	if i == 0 {
		a := list.FirstNode

		n.Next = a.Next

		tmpVal := a.Value
		a.Value = n.Value
		n.Value = tmpVal

		a.Next = &n
		return nil
	}

	// insert to i
	a, err := list.Get(i - 1)
	if err != nil {
		return err
	}
	b := a.Next
	a.Next = &n
	n.Next = b
	return nil
}

func (list *LinearList) Delete(i int) error {
	l := list.Length()
	if l == 0 {
		return ErrLinearListHasNoEntry
	}
	if i < 0 || i > l-1 {
		return ErrLinearListOutOfBounds
	}

	// delete first node
	if i == 0 {
		list.FirstNode = list.FirstNode.Next
		return nil
	}

	// delete node at i
	a, err := list.Get(i - 1)
	if err != nil {
		return err
	}
	b := a.Next
	a.Next = b.Next
	return nil
}

func (list *LinearList) Get(i int) (*LinearListNode, error) {
	if i < 0 || i > list.Length()-1 {
		return nil, ErrLinearListOutOfBounds
	}
	n := list.FirstNode
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

func (list *LinearList) Length() int {
	n := list.FirstNode
	if n == nil {
		return 0
	}

	l := 1
	for {
		n = n.Next
		if n == nil {
			break
		}
		l++
	}
	return l
}

func (list *LinearList) ForEach(f func(*LinearListNode)) {
	n := list.FirstNode
	for {
		if n == nil {
			break
		}
		f(n)
		n = n.Next
	}
}

func (list *LinearList) ToSlice() []*LinearListNode {
	var s []*LinearListNode
	list.ForEach(func(n *LinearListNode) {
		s = append(s, n)
	})
	return s
}
