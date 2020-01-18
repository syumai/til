package add

import "testing"

func TestAdd(t *testing.T) {
	result, err := Add(1, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 3 {
		t.Errorf("want: 3, got: %d", result)
	}
}
