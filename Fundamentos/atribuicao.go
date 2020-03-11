package main

import "log"

func main(){
  var b = 3
  log.Println(b)
  
  i := 3
  i += 3
  i -= 3
  i /= 2
  i *= 2
  i %= 2
  
  log.Println(i)
  
  x, y := 1, 2
  log.Println(x, y)
  
  x, y = y, x
  log.Println(x, y)
}
