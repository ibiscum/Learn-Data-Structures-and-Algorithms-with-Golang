// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing testing package
import (
	"testing"
)

// NewKnowledgeGraph test method
func TestNewKnowledgeGraph(test *testing.T) {

	// var knowledgeGraph *KnowledgeGraph

	var knowledgeGraph = NewKnowledgeGraph()

	test.Log(knowledgeGraph)

	if knowledgeGraph == nil {

		test.Errorf("error in creating a knowledgeGraph")
	}

}
