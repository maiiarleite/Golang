package main

import "log"

func obterResultado(nota float64) string{
  if nota >= 6{
    return "Aprovado!"
  }
  return "Reprovado!"
}

func main(){
  log.Println(obterResultado(6.2))
}
