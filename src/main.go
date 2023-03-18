package main

import "fmt" // paquete para imprimir en consola

func main() {

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	// Division
	result = y / x
	fmt.Println("Division: ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	// Incrementar
	x++
	fmt.Println("Incrementar: ", x)

	// Decrementar
	x--
	fmt.Println("Decrementar: ", x)

}
