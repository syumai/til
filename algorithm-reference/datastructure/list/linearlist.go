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
	if i < 0 {
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
	if list.Blank() {
		return ErrLinearListHasNoEntry
	}
	if i < 0 {
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
	if b == nil {
		return ErrLinearListOutOfBounds
	}

	a.Next = b.Next
	return nil
}

func (list *LinearList) Get(i int) (*LinearListNode, error) {
	if list.Blank() {
		return nil, ErrLinearListHasNoEntry
	}
	if i < 0 {
		return nil, ErrLinearListOutOfBounds
	}

	n := list.FirstNode
	j := 0
	for {
		if i == j {
			break
		}

		n = n.Next
		if n == nil {
			return nil, ErrLinearListOutOfBounds
		}
		j++
	}
	return n, nil
}

func (list *LinearList) Blank() bool {
	return list == nil || list.FirstNode == nil
}

func (list *LinearList) Length() int {
	if list.Blank() {
		return 0
	}

	n := list.FirstNode
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

func (list *LinearList) Range(f func(int, *LinearListNode)) {
	i := 0
	n := list.FirstNode
	for {
		if n == nil {
			break
		}
		f(i, n)
		i++
		n = n.Next
	}
}

func (list *LinearList) ToSlice() []*LinearListNode {
	var s []*LinearListNode
	list.Range(func(_ int, n *LinearListNode) {
		s = append(s, n)
	})
	return s
}
