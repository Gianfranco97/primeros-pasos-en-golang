package main

import "fmt"

func main() {
	var name string

	fmt.Println("¿Cómo te llamas?")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hola %s", name)
}
