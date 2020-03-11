package main

import (
	"log"
	"fmt"
)

func f1(){
	log.Println("F1")
}

func f2(p1 string, p2 string){
	log.Printf("F2: %s %s\n", p1, p2)
}

func f3() string{
	return "F3"
}

func f4(p1, p2 string) string{
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func  f5() (string, string)  {
	return "Retorno 1", "Retorno 2"
}

func main(){
	f1()
	f2("Param1", "Param2")

	r2, r3 := f3(), f4("Param1", "Param2")
	log.Println(r2)
	log.Println(r3)

	r51, r52 := f5()
	log.Println("F5:", r51, r52)
}