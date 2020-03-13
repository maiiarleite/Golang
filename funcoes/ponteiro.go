package main

import "log"

func inc1(n int) {
	n++
}

func inc2(n*int) {
	*n++
}

func main() {
	n := 1

	inc1(n)
	log.Println(n)

	inc2(&n)
	log.Println(n)
}