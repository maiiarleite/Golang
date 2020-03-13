package main

import "log"

func main(){
	funcESalarios := map[string]float64{
		"José Benedito": 12000.98,
		"Florentina da Silva": 40000.56,
		"Clementina Ferreira": 3000.9,
	}

	funcESalarios["Pedro Filho"] = 2400.79
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios {
		log.Println(nome, salario)
	}
}