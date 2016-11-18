package grains

import (
	"errors"
	"math"
)

const testVersion = 1
func Square(s int) (uint64,error) {
	var result int
	if s < 1 || s > 64 {
		return 0, errors.New("Not a valid square")
	}
	for num:=1; num <= s; num++ {
		if num == 1 {
			result = 1
		} else {
			result = result * 2
		}

	}
	return uint64(result),nil
}
/*
func Total() uint64{
	var result uint64
	var s = uint64(65)
	for num:=uint64(1); num <= s; num++ {
		if num == uint64(1) {
			result = uint64(1)
		} else {
			result += result * uint64(2)
		}

	}
	return result
}
*/
func Total() uint64 {
	return uint64(math.Pow(2, float64(65))) - 1
}
