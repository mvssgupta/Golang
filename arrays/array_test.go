package main

import "testing"

func TestArraySum(t *testing.T) {

	//we can initialize an array in the following two ways:

	//[N]type{value1, value2, ..., valueN} e.g. nums := [3]int{1, 2, 3}
	// [...]type{value1, value2, ..., valueN} e.g. nums := [...]int{1, 2, 3, 4, 5}

	nums := []int{1, 2, 3, 4}
	got := sum(nums)
	got2 := sumRange(nums)
	expected := 10

	if got != expected && got2 != expected {
		t.Errorf("got %d when expected %d, given array %v", got, expected, nums)
	}
}
