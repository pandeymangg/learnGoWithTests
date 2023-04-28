package main

import (
	"fmt"

	arraysandslices "github.com/pandeymangg/learnGoWithTests/arraysAndSlices"
	"github.com/pandeymangg/learnGoWithTests/helloWorld"
	"github.com/pandeymangg/learnGoWithTests/integers"
	"github.com/pandeymangg/learnGoWithTests/iteration"
)

func main() {
	fmt.Println(
		helloWorld.Hello("Anshuman", "Hindi"))
	fmt.Println(integers.Add(2, 20))

	fmt.Println(iteration.Repeat("c", 69))

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2}

	sliceSum := arraysandslices.SumAllTails(slice1, slice2)

	fmt.Println(sliceSum)

}
