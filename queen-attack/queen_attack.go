package queenattack

import (
	"regexp"
	"errors"
	"math"
)

const testVersion = 1

func CanQueenAttack (q1, q2 string) (bool, error) { //canAttack, isValid
var canAttack bool
var isValid error

isValid = validColumnRow(q1,q2)

	if isValid == nil {
		rowdst:=math.Abs(float64(int(q1[0]) - int(q2[0])))
		coldst:=math.Abs(float64(int(q1[1]) - int(q2[1])))
		if rowdst == coldst || rowdst==0 || coldst==0 {
			canAttack = true
		}

	}

return canAttack, isValid
}

func validColumnRow(q1,q2 string) error {
	re := regexp.MustCompile("[a-h][1-8]")
	if re.FindStringIndex(q1)  == nil || re.FindStringIndex(q2) == nil || q1 == q2 {
		return errors.New(q1 + " or " + q2 + " location is not valid")
	}
return nil
}
