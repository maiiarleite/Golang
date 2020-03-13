package main

import (
	"log"
	"time"
)

func doisTresQuatroVezes(base int, ch chan int){
	time.Sleep(time.Second)
	ch <- 2 * base

	time.Sleep(time.Second)
	ch <- 3 * base

	time.Sleep(time.Second)
	ch <- 4 * base
}

func main() {
	ch := make(chan int, 3)
	go doisTresQuatroVezes(2, ch)

	a, b:= <-ch, <-ch
	log.Println(a, b)

	log.Println(<-ch)
}