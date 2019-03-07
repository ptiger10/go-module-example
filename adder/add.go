package adder

// Add an arbitrary collection of ints
// Return an int
// Not that interesting
// Another try
func Add(nums ...int) int {
	var answer int
	for _, num := range nums {
		answer += num
	}
	return answer
}
