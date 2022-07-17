package main

import "fmt"

func Hello(name string) string {
	const englishHelloPrefix = "Hello, "
	// return "Hello, world"
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("jack"))
}
