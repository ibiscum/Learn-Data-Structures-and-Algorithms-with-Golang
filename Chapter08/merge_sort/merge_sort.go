// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
	"math/rand"
)

// create array
func createArray(num int) []int {
	// var array []int
	var array = make([]int, num)
	// rand.Seed(time.Now().UnixNano())
	var i int
	for i = 0; i < num; i++ {
		array[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return array
}

// MergeSorter algorithm
func MergeSorter(array []int) []int {

	if len(array) < 2 {
		return array
	}
	// var middle int
	var middle = (len(array)) / 2
	return JoinArrays(MergeSorter(array[:middle]), MergeSorter(array[middle:]))
}

// Join Arrays method
func JoinArrays(leftArr []int, rightArr []int) []int {

	var num int
	var i int
	var j int
	num, i, j = len(leftArr)+len(rightArr), 0, 0
	// var array []int
	var array = make([]int, num)

	var k int
	for k = 0; k < num; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			array[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			array[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			array[k] = leftArr[i]
			i++
		} else {
			array[k] = rightArr[j]
			j++
		}
	}
	return array
}

// main method
func main() {

	// var elements []int
	var elements = createArray(40)
	fmt.Println("--- Before Sorting ---\n\n", elements)
	fmt.Println("\n--- After Sorting ---\n\n", MergeSorter(elements))
}
