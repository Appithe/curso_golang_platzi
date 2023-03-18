package main

import "fmt" // paquete para imprimir en consola

func main() {
	// $go build #para compilar el código y generar el archivo build
	// $go run #para ejecutar el código sin compilar

	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaración de variables enteras
	base := 12          // variable no declarada anteriormente
	var altura int = 14 // declaramos con tipo de dato y asignamos un valor
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int     // 0
	var b float64 // 0
	var c string  // " "
	var d bool    // false

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)

}
