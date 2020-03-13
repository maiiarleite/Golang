package main

import "log"

func imprimirAprovados(aprovados ...string){
	log.Println("Lista de Aprovados:")
	for i, aprovado := range aprovados{
		log.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...)
}
