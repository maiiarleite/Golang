package main

import "log"

func fatorial(n uint) uint{
	switch {
	case n == 0:
		return 1
	default:
		return n*fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	log.Println(resultado)
}