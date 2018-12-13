package main

// Sum combines numbers of an array
// and fixes size of array (5)
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum

}

// Sum2 is same as Sum, but different layout
// and is a slice (allows any size array)
func Sum2(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {}
