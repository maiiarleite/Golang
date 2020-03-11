package main

import "log"

type carro struct {
	nome string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	log.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)
	log.Println(f)
}
