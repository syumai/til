package gomod117

func ConvertSliceToArrayOf5Ints(s []int) [5]int {
	p := (*[5]int)(s)
	return *p
}
