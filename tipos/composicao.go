package main

import "log"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bwm7 struct {}

func (b bwm7) ligarTurbo() {
	log.Println("Turbo...")
}

func (b bwm7) fazerBaliza() {
	log.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()
}