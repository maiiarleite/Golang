package main

import (
	"log"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	log.Println("Só depois que foi lido")
}

func main(){
	c := make(chan int, 6)

	c <- 5
	c <- 6

	go rotina(c)

	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
}