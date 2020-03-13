package main

import "log"

func main(){
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 74764.76,
			"Gustavo Pereira": 98723.98,
		},
		"J": {
			"Joao Marcos": 7623.65,
		},
		"M": {
			"Marcelo Filho": 21658.43,
		},
	}

	delete(funcsPorLetra,"J")

	for letra, funcs := range funcsPorLetra {
		log.Println(letra, funcs)
	}
}