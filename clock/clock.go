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

	if hour >= 24 {
		hrs = hour - 24 * (hour/24)
	}
	if hour < 0 {
		hrs = 24 - (24 * ( hour/24) - hour)
	}
	if minute >= 60 {
		hrs = hrs + (minute/60)
		if hrs >= 24 {
			hrs = hrs - 24 * (hrs/24)
		}
		mns = minute - 60 * (minute/60)
	}
	return Clock{h:hrs,m:mns}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m + minutes)
}

