package main

import "fmt"

func main() {
	fmt.Println(Hello("", ""))
}

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
	spanish       = "Spanish"
	french        = "French"
)

func Hello(n, l string) string {
	if n == "" {
		n = "World"
	}
	return greetingPrefix(l) + n
}

func greetingPrefix(l string) (prefix string) {
	switch l {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
