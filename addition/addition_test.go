package integers

import (
	"fmt"
	"testing"
)

//short hand for writing x int , y int is x,y int
func Add(x, y int) int {
	return x + y
}
func ExampleAdd() {
	sum := Add(1, 5)
	// " 'Output:' is used for telling the expected value to the testing framework of Golang"
	//if we remove that then the sum function don't have an expected value to compare so it wont gets executed.
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("got %d ,expected %d", sum, expected)
	}

}
