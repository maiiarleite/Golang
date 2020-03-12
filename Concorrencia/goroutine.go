package main

import (
	"log"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		log.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main(){
	go fale("Joao","Ok", 3)
	fale("Patricia", "Fim", 7)
}