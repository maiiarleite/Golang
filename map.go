package main

import "log"

func main(){
	aprovados := make(map[int] string)

	aprovados[12345678978] = "Maria"
	aprovados[57842847578] = "Paulo"
	aprovados[98747366263] = "João"
	log.Println(aprovados)

	for cpf, nome := range aprovados {
		log.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	log.Println(aprovados[57842847578])
	delete(aprovados, 57842847578)
	log.Println(aprovados[57842847578])
}
