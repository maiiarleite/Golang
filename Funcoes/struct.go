package main

import "log"

type produto struct {
	nome string
	preco float64
	desconto float64
}

func (p produto) precoComDesconto() float64{
	return  p.preco*(1-p.desconto)
}

func main(){
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.67,
		desconto: 0.02,
	}
	log.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1489.87, 0.43}
	log.Println(produto2.precoComDesconto())
}