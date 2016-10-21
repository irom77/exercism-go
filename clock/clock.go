// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	hrs := hour
	mns := minute

	time := (hour*60 + minute) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}

	hrs = time / 60
	mns = time % 60

	return Clock{h:hrs,m:mns}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m + minutes)
}

