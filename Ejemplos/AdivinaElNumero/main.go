package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var respuestaCorrecta = 1 + rand.Intn(10)
	var respuesta int

	fmt.Println("Intentan adivinar el número, introduce un número entre el uno y el diez...")
	for {
		fmt.Scanf("%d", &respuesta)

		if respuesta != respuestaCorrecta {
			fmt.Println("Lo siento, te equivocaste. Intenta nuevamente...")
		} else {
			fmt.Println("Felicidades, adivinaste el número :D")
			break
		}
	}
}
