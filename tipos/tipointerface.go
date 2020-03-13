package main

import "log"

type curso struct {
	nome string
}

func main()  {
	var coisa interface{}
	log.Println(coisa)

	coisa = 3
	log.Println(coisa)

	type dinamico interface {}

	var coisa2 dinamico = "Opa"
	log.Println(coisa2)

	coisa2 = true
	log.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a linguagem do Google"}
	log.Println(coisa2)
}