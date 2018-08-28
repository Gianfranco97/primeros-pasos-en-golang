package main

import (
	"fmt"
)

func main() {
	var numeroA float32
	var numeroB float32
	var total float32

	fmt.Println("A continuación se le va a solicitar dos números para sumar (A y B)...")
	fmt.Print("A: ")
	fmt.Scanf("%f", &numeroA)
	fmt.Print("B: ")
	fmt.Scanf("%f", &numeroB)

	total = numeroA + numeroB
	fmt.Printf("La suma de los números es igual a: %.2f", total)
}
