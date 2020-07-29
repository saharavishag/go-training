package main

import "fmt"

// interface type
type bot interface {
	getGreeting() string
}

// concrete type
type enBot struct{}
type spBot struct{}

func main() {
	eb := enBot{}
	sb := spBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (enBot) getGreeting() string {
	return "Hello"
}

func (spBot) getGreeting() string {
	return "Hola"
}
