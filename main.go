package main

import "fmt"

func greetings(input string, f func(string)) {
	f(input)
}

func main() {
	str := "I must disappear once"
	greetings(str, func(s string) {
		fmt.Println(s)
	})
}
