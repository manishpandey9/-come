package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
		// https://golang.org/pkg/fmt/#hdr-Printing
	}
}

/*Writing a test is just like writing a function, with a few rules

It needs to be in a file with a name like xxx_test.go
The test function must start with the word Test
The test function takes one argument only t *testing.T
In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

Need to Explore more on testing modules like why *testing.T what all other attributes
*testing have other than T
*/

/*One of the feature of go is you can launch the docs locally use
"godoc -http :8000"  to run the server  than call "http://localhost:8000/pkg"
to check all the packages available or go to "http://localhost:8000/pkg/testing/"
to check about what available in testing */
