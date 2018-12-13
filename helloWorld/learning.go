package main

import "fmt"

const defaultHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const swahiliHelloPrefix = "Jambo, "
const french = "French"
const spanish = "Spanish"
const swahili = "Swahili"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
  switch language {
  case french:
    prefix = frenchHelloPrefix
  case spanish:
    prefix = spanishHelloPrefix
  case swahili:
    prefix = swahiliHelloPrefix
  default:
    prefix = defaultHelloPrefix
  }
  return
}

func main() {
	fmt.Println(Hello("Leon", ""))
}

// As an air traffic controller
// So I can get passengers to a destination
// I want to instruct a plane to land at an airport
