package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		//semelhante ao length
		go func(url string){
			resp, _ := http.Get(url)
			//chama a estrutura da url
			html, _ := ioutil.ReadAll(resp.Body)
			//ioutil - leitura (manipulacao de buffer)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			//expressao regular - autômato
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://pt.wikipedia.org/wiki/Wikip%C3%A9dia:P%C3%A1gina_principal", "https://www.youtube.com")
	log.Println("Primeiros:", <-t1, "|", <-t2)
	log.Println("Segundos:", <-t1, "|", <-t2)
}