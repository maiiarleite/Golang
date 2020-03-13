package main

import (
  "log"
  "time"
)

func main(){
  i:=1
  for i<=10{
    log.Println(i)
    i++
  }
  
  for i:=0;i<=20;i+=2{
    log.Printf("%d", i)
  }
  
  log.Println("\nMisturando...")
  for i:=1;i<=10;i++{
    if i%2==0{
      log.Print("Par ")
    }else{
      log.Print("Impar ")
    }
  }
  
  for{
    log.Println("Para sempre...")
    time.Sleep(time.Second)
  }
}
  
  
  

