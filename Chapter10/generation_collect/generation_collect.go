// /main package has examples shown
// in Go Data Structures and algorithms book
package main

import "log"

func GetObjectsFromOldGeneration(gen int) {

}

func GenerationCollect() {

	// var currentGeneration int
	var currentGeneration = 3

	var objects *[]object
	objects = GetObjectsFromOldGeneration(currentGeneration)

	// var object *object

	for _, object := range objects {
		var markedAlready bool = true

		// markedAlready := IfMarked(object)
		if markedAlready {

			log.Printf("object: %v", object)
			//map[object] = true
		}
	}
}
