package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"
	danish  = "Danish"

	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	danishHelloPrefix  = "Hej, "
)

func Hello(name, langauge string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(langauge) + name

}

func greetingPrefix(langauge string) (prefix string) {
	switch langauge {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case danish:
		prefix = danishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
