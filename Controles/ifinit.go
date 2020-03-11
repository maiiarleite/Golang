package main

import (
  "log"
  "math/rand"
  "time"
)

func numeroAleatorio() int{
  s:=rand.NewSource(time.Now().UnixNano())
  r:=rand.New(s)
  return r.Intn(10)
}

func main(){
  if i:=numeroAleatorio(); i>5{
    log.Println("Ganhou!!!")
    log.Println("->", i)
  }else{
    log.Println("Perdeu!!!")
    log.Println("->", i)
  }
}
