package main

import "fmt"

func main() {

	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	fmt.Print("\n")

	// For while
	counter := 0
	for counter < 10 {
		fmt.Print(counter)
		counter++
	}

	fmt.Print("\n")

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
