package main

import "fmt"

func main() {
	a := 20
	b := 10

	//SUMA
	result := a + b // := sirve para declarar e inicializar
	fmt.Println("Suma:", result)

	// RESTA
	result = a - b // Se usa "=" porque la variable ya está declarada, entonces solo se está reutilizando
	fmt.Println("Resta:", result)

	// MULTIPLICACION
	result = a * b
	fmt.Println("Multiplicacion:", result)

	// DIVISION
	// La unica manera de obtener una división exacta es que los datos sean float (Ex: 2.0, 1.5,...) de otra manera se obtiene división entera
	result = a / b
	fmt.Println("Division Entera:", result)

	result2 := 2 / 3
	fmt.Println("Division Entera:", result2)

	result3 := 2.0 / 3.0
	fmt.Println("Division Exacta:", result3)

	// MODULO
	result = a % b
	fmt.Println("Modulo:", result)

	result = 5 % 2
	fmt.Println("Modulo:", result)

}
