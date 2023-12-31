package integers

import (
	"fmt"
	"testing"
)

func AddedTest(t *testing.T) {

	sum := Add(2, 2)

	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)

	fmt.Println(sum)
}
