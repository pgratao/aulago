package main

import "fmt"

func main(){

	menu := map[string]float64{
		"pizza": 40.00,
		"suco": 12.50,
		"xtudo": 22.90,
	}

	fmt.Println(menu)

	for k, v := range menu{
		fmt.Println(k, "R$", v)
	}

	contanova := novaConta("Astrubal")
	fmt.Println(contanova)
}