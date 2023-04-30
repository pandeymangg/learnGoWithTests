package main

import (
	"log"
	"net/http"

	"github.com/pandeymangg/learnGoWithTests/dependency_injection"
)

func main() {
	// fmt.Println(
	// 	helloWorld.Hello("Anshuman", "Hindi"))
	// fmt.Println(integers.Add(2, 20))

	// fmt.Println(iteration.Repeat("c", 69))

	// slice1 := []int{1, 2, 3}
	// slice2 := []int{1, 2}

	// sliceSum := arraysandslices.SumAllTails(slice1, slice2)

	// fmt.Println(sliceSum)

	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dependency_injection.Greet(w, "John")
	})))

}
