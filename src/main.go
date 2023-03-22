package main

import "fmt" // paquete para imprimir en consola

// Se usa camelCase
func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")

	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("value: ", value)

	/* value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2: ", value1, value2) */

	value1, _ := doubleReturn(2)
	fmt.Println("value1: ", value1)

	// Reto pasar a funciones el reto de calcular areas
}
