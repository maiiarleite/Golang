package main

import "log"

func obterValorAprovado(numero int) int{
	defer log.Println("fim!")
	if numero > 5000 {
		log.Println("Valor alto")
		return 5000
	}
	log.Println("Valor baixo...")
	return numero
}

func main() {
	log.Println(obterValorAprovado(6000))
	log.Println(obterValorAprovado(2000))
}