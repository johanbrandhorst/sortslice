package a

import "sort"

// IncorrectSort tries to sort an integer.
func IncorrectSort() {
	i := 5
	sortFn := func(i, j int) bool { return false }
	sort.Slice(i, sortFn) // want "sort.Slice's argument must be a slice; is called with int"
}

// CorrectSort sorts integers. It should not produce a diagnostic.
func CorrectSort() {
	s := []int{2, 3, 5, 6}
	sortFn := func(i, j int) bool { return s[i] < s[j] }
	sort.Slice(s, sortFn)
}
