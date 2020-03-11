package main

import (
  "log"
  "time"
)

func main(){
  log.Println("String: ", "Banana" == "Banana")
  log.Println("!=", 3 != 2)
  log.Println("<", 3 < 2)
  log.Println(">", 3 > 2)
  log.Println("<=", 3 <= 2)
  log.Println(">=", 3 >= 2)
  
  d1 := time.Unix(0, 0)
  d2 := time.Unix(0, 0)
  
  log.Println("Datas: ", d1 == d2)
  log.Println("Datas: ", d1.Equal(d2))
  
  type Pessoa struct{
    Nome string
  }
  
  p1 := Pessoa{"João"}
  p2 := Pessoa{"João"}
  log.Println("Pessoas: ", p1 == p2)
}
