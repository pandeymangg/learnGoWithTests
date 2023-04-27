package main

import "fmt"

const hindi = "Hindi"
const englishHelloPrefix = "Hello, "
const hindiHelloPrefix = "Namaste, "

func Hello(name string, lang string) string {
	if name == "" {
		return "Hello, World!"
	}

	if lang == hindi {
		return hindiHelloPrefix + name + "!"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("Anshuman", "Hindi"))
}
