// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing testing package
import (
	"testing"
)

// NewSocialGraph test method
func TestNewSocialGraph(test *testing.T) {

	// var socialGraph *SocialGraph

	var socialGraph = NewSocialGraph(1)

	if socialGraph == nil {

		test.Errorf("error in creating a socail Graph")
	}

}
