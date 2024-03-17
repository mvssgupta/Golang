package main

//using classic for loop

func sum(nums [4]int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return
}

//using range keyword

//here since we are using for loop to iterate an array
//we can use range keyword to iterate over it

// range lets you iterate over an array. On each iteration,
// range returns two values - the index and the value.
// We are choosing to ignore the index value by using _ blank indicator

func sumRange(nums [4]int) (result int) {
	for _, value := range nums {
		result += value
	}
	return
}
