// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
	"math/rand"
)

// randomSequence method
func randomSequence(num int) []int {

	// var sequence []int
	var sequence = make([]int, num)
	// rand.Seed(time.Now().UnixNano())
	var i int
	for i = 0; i < num; i++ {
		sequence[i] = rand.Intn(999) - rand.Intn(999)
	}
	return sequence
}

// InsertionSorter method
func InsertionSorter(elements []int) {
	var n = len(elements)
	var i int

	for i = 1; i < n; i++ {
		var j int
		j = i
		for j > 0 {
			if elements[j-1] > elements[j] {
				elements[j-1], elements[j] = elements[j], elements[j-1]
			}
			j = j - 1
		}
	}
}

// main method
func main() {

	// var sequence []int
	var sequence = randomSequence(24)
	fmt.Println("--- Before Sorting ---\n\n", sequence)
	InsertionSorter(sequence)
	fmt.Println("\n--- After Sorting ---\n\n", sequence)
}
