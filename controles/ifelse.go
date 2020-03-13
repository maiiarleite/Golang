package main

import "log"

func imprimirResultado(nota float64){
  if nota >= 7{
    log.Println("Aprovado com nota ", nota)
  }else{
    log.Println("Reprovado com nota ", nota)
  }
}

func main(){
  imprimirResultado(5.4)
  imprimirResultado(8.1)
}
