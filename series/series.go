package slice
//copied from https://github.com/patleeman/exercism-golang/blob/master/series/series.go
// All returns a list of all substrings of s with length n.

func All(n int, s string) []string {
	var substrings []string
	for i := 0; i < len(s)-n+1; i++ {
		substrings = append(substrings, s[i:i+n])
	}
	return substrings
}
//
//    // UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (string, bool) {
	var ok bool
	var ans string

	if n > len(s) || n < 0 {
		ok = false
		ans = "0"
	} else {
		ok = true
		ans = s[0:n]
	}

	return ans, ok
}
