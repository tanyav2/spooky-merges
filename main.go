package main

import "fmt"

func greetings(input string, f func(string)) {
	f(input)
}

func main() {
	str := "I must disappear once"
	if true {
		fmt.Println("I will exist")
	} else {
		greetings(str, func(s string) {
			fmt.Println(s)
		})
	}
}
