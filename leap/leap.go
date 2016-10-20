// Leap stub file

// The package name is expected by the test program.
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// take a year and report if it is a leap year
var IsLeap bool
func IsLeapYear( year int) bool {
	IsLeap = false
	if year % 4 == 0 {
		IsLeap = true
	}
	if year % 100 == 0 {
		IsLeap = false
		if year % 400 == 0 {
			IsLeap = true
		}
	}
	return IsLeap
}
