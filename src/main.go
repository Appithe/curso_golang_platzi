/* Reto
Calcular area de un circulo y un trapecio y rectángulo */

package main

import "fmt" // paquete para imprimir en consola

func areaCuadrado(lado int) int {
	return lado * lado
}

func areaRectangulo(base, altura int) int {
	return base * altura
}

func areaCirculo(radio float64) float64 {
	const PI float64 = 3.14

	return PI * radio * radio
}

func areaTrapecio(baseA, baseB, altura float64) float64 {
	return ((baseA + baseB) * altura) / 2
}

func main() {

	// Area cuadrado
	areaCuadrado := areaCuadrado(10)
	fmt.Println("Area cuadrado: ", areaCuadrado)

	// Rectángulo
	areaRectangulo := areaRectangulo(20, 10)
	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo
	areaCirculo := areaCirculo(10)
	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	areaTrapecio := areaTrapecio(6, 15, 25)
	fmt.Println("El Area del Trapecio es :", areaTrapecio)

}
