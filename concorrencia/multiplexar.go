package main

import (
	html "titulo"
	"log"
)
func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
	html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
	html.Titulo("https://pt.wikipedia.org/wiki/Wikip%C3%A9dia:P%C3%A1gina_principal", "https://www.youtube.com"),
	)
	log.Println(<-c, "|", <-c)
	log.Println(<-c, "|", <-c)
}