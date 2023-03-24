package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value, ok := m["jose"] // "ok" retorna true si la clave existe
	fmt.Println(value, ok)

	value1, ok := m["josep"] // "ok" retorna true si la clave existe
	fmt.Println(value1, ok)

}
