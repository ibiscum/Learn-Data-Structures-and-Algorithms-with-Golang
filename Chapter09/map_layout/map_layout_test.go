// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing testing package
import (
	"testing"
)

// NewMapLayout test method
func TestNewMapLayout(test *testing.T) {

	// var mapLayout *MapLayout

	var mapLayout = NewMapLayout()

	test.Log(mapLayout)

	if mapLayout == nil {

		test.Errorf("error in creating a mapLayout")
	}

}
