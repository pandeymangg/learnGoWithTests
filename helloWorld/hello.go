package helloWorld

const hindi = "Hindi"
const french = "French"

const englishHelloPrefix = "Hello, "
const hindiHelloPrefix = "Namaste, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		return "Hello, World!"
	}

	prefix := greetingPrefix(lang)

	return prefix + name + "!"
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix

	case hindi:
		prefix = hindiHelloPrefix

	default:
		prefix = englishHelloPrefix
	}

	return
}
