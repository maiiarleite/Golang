package main

import (
  "log"
  "time"
)

func main(){
  t:=time.Now()
  switch{
    case t.Hour() < 12:
      log.Println("Bom dia!")
    case t.Hour() < 18:
      log.Println("Boa tarde!")
    default:
      log.Println("Boa noite!")
  }
}




