package main

import "log"

func main(){
  i:=1
  
  var p *int = nil
  p=&i
  *p++
  i++
  
  log.Println(p, *p, i, &i)
}
