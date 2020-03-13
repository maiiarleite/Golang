package main

import "log"

func closure() func() {
	x := 10
	var funcao = func() {
		log.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	log.Println(x)

	imprimeX := closure()
	imprimeX()
}
