package levenshtein

import (
	"strings"
	"unicode/utf8"
)

func sort(s1, s2 string) (shorter, longer string) {
	if utf8.RuneCountInString(s1) > utf8.RuneCountInString(s2) {
		return s2, s1
	}
	return s1, s2
}

func minimum(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func rowVector(s1, s2 []rune) (vector []int) {
	n := minimum(len(s1), len(s2))
	x := make([]int, n+1)

	for i := 1; i <= n; i++ {
		x[i] = i
	}
	return x
}

// Distance calculates levenshtein distance of 2 strings
// the function changes the strings to lowercase
func Distance(a, b string) int {
	s1, s2 := sort(strings.ToLower(a), strings.ToLower(b))

	// If s1 = s2, then the levenshtein distance is 0
	if s1 == s2 {
		return 0
	}

	// Returns to the longest string value if s1 is Null
	if utf8.RuneCountInString(s1) == 0 {
		return utf8.RuneCountInString(s2)
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	x := rowVector(r1, r2)

	for i := 1; i <= len(r2); i++ {
		prev := i
		var current int
		for j := 1; j <= len(r1); j++ {
			if r2[i-1] == r1[j-1] {
				current = x[j-1] // match
			} else {
				current = minimum(minimum(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[len(r1)] = prev
	}
	return x[len(r1)]
}
