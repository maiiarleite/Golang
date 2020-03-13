package main

import "log"

func main() {
	var soma = func(a, b int) int {
		return a+b
	}

	log.Println(soma(4, 7))

	sub := func(a, b int) int {
		return a-b
	}

	log.Println(sub(2, 6))
}