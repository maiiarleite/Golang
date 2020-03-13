package main

import (
	"log"
	"fmt"
)
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _:= fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	log.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		log.Println(err)
	}
}