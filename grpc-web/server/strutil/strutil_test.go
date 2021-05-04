package strutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s string
		want string
	}{
		{
			name: "Valid",
			s: "example",
			want: "elpmaxe",
		},
		{
			name: "Valid blank",
			s: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
