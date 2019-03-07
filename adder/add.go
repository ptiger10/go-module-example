package adder

// Add an arbitrary collection of ints
func Add(nums ...int) int {
	var answer int
	for _, num := range nums {
		answer += num
	}
	return answer
}
