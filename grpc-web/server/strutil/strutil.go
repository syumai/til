package strutil

// Reverse reverses string
func Reverse(s string) string {
	s2 := make([]rune, len(s))
	l := len(s)
	for i, r := range s {
		s2[l-i-1] = r
	}
	return string(s2)
}
