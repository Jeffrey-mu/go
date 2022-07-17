package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	// return "Hello, world"
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("jack"))
}
