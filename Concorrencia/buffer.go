package main

import (
	"log"
	"time"
)

func rotina(ch chan int){
	time.Sleep(time.Second)
	//println aleatorio
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	log.Println("Executou!")
	ch <- 5
	ch <- 6
}

func main(){
	ch := make(chan int, 3)

	go rotina(ch)
	//time.Sleep(time.Second)
	//println aleatorio
	log.Println(<-ch)
	log.Println(<-ch)
	log.Println(<-ch)
	log.Println(<-ch)
	log.Println(<-ch)
	log.Println(<-ch)
}
