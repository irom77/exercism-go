package diffsquares

func SquareOfSums (n int) int {
	var result int
	var s = iter(n)
	for _, n := range s {
		result = result + n
	}
 return result * result
}

func SumOfSquares (n int) int {
	var result int
	var s = iter(n)
	for _, n := range s {
		result += n * n
	}
	return result
}

func Difference (n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func iter (n int ) []int {
	var s []int
	for i := 1; i <= n; i++ {
		//fmt.Println(i)
		s = append(s,i)
	}
	return s
}