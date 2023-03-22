package main

import (
	"fmt"
)

func main() {
	// Defer --> antes de terminar de ejecutarse la función ejecuta el código que le sigue en la linea, se recomienda usar una por función; se puede usar para cerrar sesiones en DDBB
	defer fmt.Print("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue --> se puede usar para manejo de errores
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

}
