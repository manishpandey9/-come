package arrayandslices

func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	//range lets you iterate over an array. On each iteration, range returns two values - the index and the value.
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	return sum
}

/* Arrays and their type
size of an array is encoded in it's type If you try to pass an [4]int into a function that expects [5]int, it won't compile. They are different types so it's just the same as trying to pass a string into a function that wants an int.
*/
