package benchmark

func Bubble(input []int) []int {
	// n is the number of items in our list
	// set swapped to true
	swapped := true
	// loop
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in our list
		for i := 1; i < len(input); i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i-1] > input[i] {
				// swap values using Go's tuple assignment
				input[i], input[i-1] = input[i-1], input[i]
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}
	// finally, print out the sorted list
	return input
}
