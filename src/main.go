package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	v1 := 1
	v2 := 2

	if v1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With AND
	if v1 == 1 && v2 == 2 {
		fmt.Println("Es verdad")
	}

	// With OR
	if v1 == 0 || v2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	// convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil { // nil --> es true cuando una función determinada tiene un error
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

	/*
		Retos:
		1.- Determinar si un numero es par o impar
		2.- Validar usuario y contraseña
	*/
}
