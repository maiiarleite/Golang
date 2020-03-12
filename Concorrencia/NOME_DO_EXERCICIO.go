package main

import (
	"log"
	"time"
	"fmt"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3 ; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("\n%s falando %d", pessoa, i)
		}
	}()
	return c
}

func juntando(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := juntando(falar("João"), falar("Maria"))
	log.Println(<-c, <-c)
	log.Println(<-c, <-c)
	log.Println(<-c, <-c)
}
