package main

import (
	"encoding/json"
	"log"
)

type produto struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
	Tags []string `json:"tags"`
}

func main() {
	p1 := produto{1, "Notebook", 2000.00, []string{"Promoção","Eletrônicos"}}
	p1Json, _ := json.Marshal(p1)
	log.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id":2,"nome": "Lapis Branco", "preco": 2.69, "tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	log.Println(p2.Tags[1])
}
