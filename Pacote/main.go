package main

import (
	goarea "area"
	"log"
	//criar novo projeto: area | arquivo Go: goarea
)

func main() {
	log.Println(goarea.Circulo(2.0))
	log.Println(goarea.Retangulo(4.0, 2.0))
	log.Println(goarea.Triangulo(4.0, 3.0))
}
