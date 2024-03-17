package loops

// To do stuff repeatedly in Go, you'll need 'for'. In Go there are no while, do while loops, you can only use 'for'

import (
	"fmt"
	"testing"
)

//instead of writing 'return result' , we can write the variable name
// at the start itself along with the return type and we can skip at the end
//this process is called 'named return value'.

//to pass default values in Go , we need to use
//variadic functions which means variable parameter size using ... operator
//and directly assign a default value in the function body

func Iterate(char string, nums ...int) (result string) {
	num := 5           //assigning a default value
	if len(nums) > 0 { //overrding the default value if arguments are passed to the function
		num = nums[0]
	}
	i := 1
	for i <= num {
		result += char
		i++
	}
	return
}

func ExampleIterate() {
	result := Iterate("a", 3)
	fmt.Println(result)
	//Output:aaa
}

func TestLoop(t *testing.T) {
	repeated := Iterate("a")
	expected := "aaaaa"

	if repeated != expected {
		//we are using %q since it helps to highlight the variable with double quotes
		t.Errorf("got %q expected %q", repeated, expected)
	}
}
