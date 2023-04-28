package main

import (
	"fmt"

	"github.com/pandeymangg/learnGoWithTests/helloWorld"
	"github.com/pandeymangg/learnGoWithTests/integers"
	"github.com/pandeymangg/learnGoWithTests/iteration"
)

func main() {
	fmt.Println(
		helloWorld.Hello("Anshuman", "Hindi"))
	fmt.Println(integers.Add(2, 20))

	fmt.Println(iteration.Repeat("c", 69))

}
