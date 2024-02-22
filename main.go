package main

import (
	"fmt"

	"github.com/rocha7778/frutas/pojos"
)

func main() {

	fmt.Println("Hello World!")
	f := pojos.Fruta{
		Name:  "Banana",
		Color: "Azul",
		Value: 1.5,
	}
	fmt.Println(f)
}
