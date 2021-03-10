package main

import "fmt"

type bot interface {
	// 如果你有一个方法叫getGreeting()
	// 而且返回了一个string 那你就是一种bot
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// some logic
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
