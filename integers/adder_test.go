package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// here if you will not add the above comment
//the code will get compiled but not excute

// go test -v   (using this command you can see which all functions are running together)

// https://getstream.io/blog/how-a-go-program-compiles-down-to-machine-code/   follow this link to deepley understand how to go code is compiled into machine code
