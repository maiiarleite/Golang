package main

import (
	"log"
	"math"
)

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Distancia(a, b Ponto) float64{
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

func main () {
	p1 := Ponto{2.0, 3.0}
	p2 := Ponto{3.0, 3.0}

	log.Println(catetos(p1, p2))
	log.Println(Distancia(p1, p2))
}