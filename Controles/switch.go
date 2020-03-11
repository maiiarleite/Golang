package main

import "log"
  
func notaParaConceito(n float64) string{
  var nota = int(n)
  
  switch nota{
    case 10:
      fallthrough
    case 9:
      return "A"
    case 8, 7:
      return "B"
    case 6, 5:
      return "C"
    case 4, 3:
      return "D"
    case 2, 1, 0:
      return "E"
    default:
    return "Nota inválida"
  }
}

func main(){
  log.Println(notaParaConceito(9.4))
  log.Println(notaParaConceito(6.9))
  log.Println(notaParaConceito(2.5))
}




