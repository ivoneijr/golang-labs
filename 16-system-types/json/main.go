package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct to json

	p1 := product{1, "Macbook", 17000.00, []string{"Promo", "apple"}}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	var p2 product
	jsonString := `{"id":1,"name":"Macbook","price":17000,"tags":["Promo","apple"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
