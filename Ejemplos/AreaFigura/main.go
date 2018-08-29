package main

import "fmt"

type Rectangulo struct {
	base, altura float32
}

type Trapecio struct {
	baseMayor, baseMenor, altura float32
}

func (figura Rectangulo) area() float32 {
	return figura.base * figura.altura
}

func (figura Trapecio) area() float32 {
	return (figura.baseMayor + figura.baseMenor) * figura.altura / 2
}

func main() {
	rectangulo := Rectangulo{base: 4, altura: 7.5}
	trapecio := Trapecio{baseMayor: 5, baseMenor: 2, altura: 3}
	fmt.Printf("Area del rectángulo: %f\n", rectangulo.area())
	fmt.Printf("Area del trapecio: %f\n", trapecio.area())
}
