package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"` // identificador com letra maiuscula e publica, minuscula e privado
	Preco float64  `json:"preco"`
	Tags  []string `json: "tags"`
}

func main() {

	// Struct pra JSON
	p1 := produto{1, "Notebook", 7800.00, []string{"Promoção", "Eletrônico"}}
	p1ToJson, _ := json.Marshal(p1)

	fmt.Println(string(p1ToJson))

	// JSON to struct
	var p2 produto
	jsonString := `{"id":1,"nome":"Notebook","preco":7800,"Tags":["Promoção","Eletrônico"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)

}
