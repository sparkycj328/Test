package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

// Creates a constant to cut down on creating string every time
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}
