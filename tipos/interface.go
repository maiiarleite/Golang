package main

import (
	"log"
	"fmt"
)

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome string
	sobrenome string
}

type produto struct {
	nome string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	log.Println(x.toString())
}

func main (){
	var coisa imprimivel = pessoa{"Carlos", "Filho"}
	log.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 59.90}
	log.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça", 100.00}
	imprimir(p2)
}

