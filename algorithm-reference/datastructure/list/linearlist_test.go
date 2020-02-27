package list

import (
	"fmt"
	"reflect"
	"testing"
)

func newLinearList(length int) *LinearList {
	if length < 0 {
		return &LinearList{}
	}
	first := &LinearListNode{
		Value: 0,
	}
	prev := first
	for i := 1; i < length; i++ {
		n := &LinearListNode{
			Value: i,
		}
		prev.Next = n
		prev = n
	}
	return &LinearList{
		FirstNode: first,
	}
}

func TestLinearList_Add(t *testing.T) {
	t.Run("ErrOutOfBounds", func(t *testing.T) {
		const linearListLength = 5
		tests := []struct {
			i     int
			valid bool
		}{
			{
				i:     -1,
				valid: false,
			},
			{
				i:     0,
				valid: true,
			},
			{
				i:     linearListLength,
				valid: true,
			},
			{
				i:     linearListLength + 1,
				valid: false,
			},
		}
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%d-%v", tt.i, tt.valid), func(t *testing.T) {
				l := newLinearList(linearListLength)
				err := l.Add(tt.i, 1)
				if tt.valid && err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if !tt.valid && err != ErrLinearListOutOfBounds {
					t.Fatalf("want err: %v, got: %v", ErrLinearListOutOfBounds, err)
				}
			})
		}
	})

	t.Run("ValidCases", func(t *testing.T) {
		const linearListLength = 2
		const valueToAdd = 999
		tests := []struct {
			name string
			i    int
			want []int
		}{
			{
				name: "Insert to top",
				i:    0,
				want: []int{valueToAdd, 0, 1},
			},
			{
				name: "Insert to middle",
				i:    1,
				want: []int{0, valueToAdd, 1},
			},
			{
				name: "Insert to last",
				i:    2,
				want: []int{0, 1, valueToAdd},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := newLinearList(linearListLength)
				err := l.Add(tt.i, valueToAdd)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				var got []int
				l.ForEach(func(n *LinearListNode) {
					got = append(got, n.Value)
				})
				if !reflect.DeepEqual(tt.want, got) {
					t.Fatalf("want: %v, got: %v", tt.want, got)
				}
			})
		}

		t.Run("Insert to blank list", func(t *testing.T) {
			l := &LinearList{}
			err := l.Add(0, 0)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			var got []int
			l.ForEach(func(n *LinearListNode) {
				got = append(got, n.Value)
			})
			want := []int{0}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("want: %v, got: %v", want, got)
			}
		})
	})
}

func TestLinearList_Delete(t *testing.T) {
	t.Run("ErrOutOfBounds", func(t *testing.T) {
		const linearListLength = 5
		tests := []struct {
			i     int
			valid bool
		}{
			{
				i:     -1,
				valid: false,
			},
			{
				i:     0,
				valid: true,
			},
			{
				i:     linearListLength - 1,
				valid: true,
			},
			{
				i:     linearListLength,
				valid: false,
			},
		}
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%d-%v", tt.i, tt.valid), func(t *testing.T) {
				l := newLinearList(linearListLength)
				err := l.Delete(tt.i)
				if tt.valid && err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if !tt.valid && err != ErrLinearListOutOfBounds {
					t.Fatalf("want err: %v, got: %v", ErrLinearListOutOfBounds, err)
				}
			})
		}
	})

	t.Run("ErrHasNoEntry", func(t *testing.T) {
		l := &LinearList{}
		err := l.Delete(0)
		if err != ErrLinearListHasNoEntry {
			t.Fatalf("want err: %v, got: %v", ErrLinearListHasNoEntry, err)
		}
	})

	t.Run("ValidCases", func(t *testing.T) {
		const linearListLength = 3
		tests := []struct {
			name string
			i    int
			want []int
		}{
			{
				name: "Delete top",
				i:    0,
				want: []int{1, 2},
			},
			{
				name: "Delete middle",
				i:    1,
				want: []int{0, 2},
			},
			{
				name: "Delete last",
				i:    2,
				want: []int{0, 1},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := newLinearList(linearListLength)
				err := l.Delete(tt.i)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				var got []int
				l.ForEach(func(n *LinearListNode) {
					got = append(got, n.Value)
				})
				if !reflect.DeepEqual(tt.want, got) {
					t.Fatalf("want: %v, got: %v", tt.want, got)
				}
			})
		}
	})
}

func TestLinearList_Get(t *testing.T) {
	t.Run("ErrOutOfBounds", func(t *testing.T) {
		const linearListLength = 5
		tests := []struct {
			i     int
			valid bool
		}{
			{
				i:     -1,
				valid: false,
			},
			{
				i:     0,
				valid: true,
			},
			{
				i:     linearListLength - 1,
				valid: true,
			},
			{
				i:     linearListLength,
				valid: false,
			},
		}
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%d-%v", tt.i, tt.valid), func(t *testing.T) {
				l := newLinearList(linearListLength)
				n, err := l.Get(tt.i)
				if tt.valid {
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}
					if n == nil {
						t.Fatal("node must exist")
					}
					return
				}
				if !tt.valid && err != ErrLinearListOutOfBounds {
					t.Fatalf("want err: %v, got: %v", ErrLinearListOutOfBounds, err)
				}
			})
		}
	})

	t.Run("Valid", func(t *testing.T) {
		const linearListLength = 3
		l := newLinearList(linearListLength)
		n, err := l.Get(linearListLength - 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if n == nil {
			t.Fatal("node must exist")
		}
		if n.Value != linearListLength-1 {
			t.Fatalf("want node value: %d, got: %d", linearListLength-1, n.Value)
		}
	})
}

func TestListLinearList_Length(t *testing.T) {
	tests := []struct {
		name       string
		linearList *LinearList
		want       int
	}{
		{
			name: "Valid nil first node",
			linearList: &LinearList{
				FirstNode: nil,
			},
			want: 0,
		},
		{
			name:       "Valid single node",
			linearList: newLinearList(1),
			want:       1,
		},
		{
			name:       "Valid three nodes",
			linearList: newLinearList(3),
			want:       3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.linearList.Length() != tt.want {
				t.Fatalf("want: %d, got %d", tt.want, tt.linearList.Length())
			}
		})
	}
}

func TestListLinearList_ForEach_ToSlice(t *testing.T) {
	l := newLinearList(3)
	want := []int{0, 1, 2}
	var got []int
	l.ForEach(func(n *LinearListNode) {
		got = append(got, n.Value)
	})
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}
