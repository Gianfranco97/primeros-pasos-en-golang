package main

import (
	"fmt"
)

func main() {
	var tipoOperacion int8
	fmt.Println("Esta es una aplicación muy básica que solo permite realizar una operación matemática entre dos números distintos....")

	for {
		fmt.Println("¿Qué tipo de operación desea efectuar?")
		fmt.Println("1. Suma")
		fmt.Println("2. Resta")
		fmt.Println("3. multiplicación")
		fmt.Println("4. División")

		fmt.Scanf("%d", &tipoOperacion)

		if tipoOperacion > 0 && tipoOperacion < 5 {
			break
		}
		fmt.Println("Disculpe, el número introducido no es válido")
	}

	fmt.Printf("Su resultado es: %.2f", executeOperation(tipoOperacion, getData()))
}

func getData() [2]float32 {
	var datos [2]float32
	fmt.Println("Introduzca dos datos (A y B), tenga en cuenta que en la resta el dato 'A' se restará el dato 'B' y que en el caso la división el dato 'B' será el denominador del dato 'A'")
	fmt.Print("A: ")
	fmt.Scanf("%f", &datos[0])
	fmt.Print("B: ")
	fmt.Scanf("%f", &datos[1])

	return datos
}

func executeOperation(tipoOperacion int8, datos [2]float32) float32 {
	var total float32

	switch tipoOperacion {
	case 1:
		total = datos[0] + datos[1]
	case 2:
		total = datos[0] - datos[1]
	case 3:
		total = datos[0] * datos[1]
	case 4:
		total = datos[0] / datos[1]
	}

	return total
}
