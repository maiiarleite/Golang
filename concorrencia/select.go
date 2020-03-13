package main

import (
	html "titulo"
	"log"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <- c3:
		return t3
	case <-time.After(time.Millisecond * 1000):
		return "Errou!!!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://br.tradingview.com/chart/5e5e4xtA/",
		"https://www.youtube.com",
		"https://brasilescola.uol.com.br/",
	)
	log.Println(campeao)
}
