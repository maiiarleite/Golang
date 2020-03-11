package main

import "log"

func media(numeros ...float64) float64{
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total/float64((len(numeros)))
}

func main() {
	log.Printf("Media: %.2f", media(7.4, 6.3, 8.3, 9.0))
}