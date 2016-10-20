// This is a "stub" file.  It's a little start on your solution.


// The package name is expected by the test program.
// It has to stay just the way it is.
package hello

import (
	//"fmt"
)

// testVersion identifies the version of the test program that you are
// writing your code to. If the test program changes in the future --
// after you have posted this code to the Exercism site -- nitpickers
// will see that your code can't necessarily be expected to pass the
// current test suite because it was written to an earlier test version.
const testVersion = 2

//Write a function that greets the user by name, or by saying "Hello, World!" if no name is given.
//"Hello, World!" is the traditional first program for beginning programming in a new language.
func HelloWorld(name string) string {
	// Write some code here to pass the test suite.

	if name != "" {
		return "Hello, " + name + "!"
	} else {
		return "Hello, World!"
	}

}
