package main

import "fmt" // paquete para imprimir en consola

func main() {
	// declaración de variables

	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos) // %s para strings y %d para enteros
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos) // %v si no sabemos que tipo de dato estará ahi

	// Sprintf --> guarda el mensaje en una variable para después imprimirlo
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de dato --> typeof en JS
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T", cursos)

}
