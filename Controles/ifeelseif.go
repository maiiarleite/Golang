package main

import "log"

func notaParaConceito(n float64) string{
  if n >= 9 && n <= 10{
    return "A"
  }else if n >= 8 && n < 9{
    return "B"
  }else if n >= 5 && n < 8{
    return "C"
  }else if n >= 3 && n < 5{
    return "D"
  }else{
    return "E"
  }
}

func main(){
  log.Println(notaParaConceito(9.8))
  log.Println(notaParaConceito(6.9))
  log.Println(notaParaConceito(3.5))
}
