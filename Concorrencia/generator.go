package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string, 2)
	for _, url := range urls {
		go func(url string){
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("(.*?)")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")
	log.Println("Primeiros:", <-t1, "|", <-t2)
	log.Println("Segundos:", <-t1, "|", <-t2)
}