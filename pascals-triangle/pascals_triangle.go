package pascal

func Triangle (n int) [][]int {
	var result = make([][]int,n)
	for i:=0; i<n; i++ {
		result[i] = make([]int, i+1)
		for j:=0; j<=i; j++ {
			result[i][j] = Combination(i,j)
		}
	}
return result

}
//from https://github.com/btracey/math/blob/master/math.go
// Combination computes n choose k
// equal to (n!)/((n-k)! k!)
// Panics on integer overflow
// Returns 0 if either n or k are negative
func Combination(n, k int) int {
	if n < 0 || k < 0 || k > n {
		return 0
	}
	if n == 0 && k == 0 {
		return 1
	}

	// Make k the greater of the two
	if k < n-k {
		k = n - k
	}
	// Compute (n!)/(k)!, checking for overflow
	prev := 1
	mul := 1
	for i := k + 1; i <= n; i++ {
		prev = mul
		mul *= i
		if mul < prev {
			panic("math: integer overflow")
		}
	}
	// Divide by (n-k)!
	for i := 2; i <= n-k; i++ {
		mul = mul / i
	}
	return mul
}