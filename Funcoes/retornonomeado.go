package main

import "log"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

func main() {
	x, y := trocar(2, 3)
	log.Println(x, y)

	primeiro, segundo := trocar(7, 1)
	log.Println(segundo, primeiro)
}
