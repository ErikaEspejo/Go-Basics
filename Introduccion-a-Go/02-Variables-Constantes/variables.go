package main

import "fmt"

func main() {

	// VARIABLES -> Declaración e inicialización

	var nombre1 string          // Se declara
	nombre1 = "Erika"           // Se inicializa
	var nombre2 string = "Alex" // Se puede declarar e inicializar a la vez
	nombre3 := "Juan"           // Nos evitamos especificar variable y tipo de dato. Go detecta el tipo de dato.

	fmt.Println(nombre1, nombre2, nombre3)

	// VARIABLES -> Declaradas sin inicializar, de esta manera tienen un valor por defecto

	var a int     // Default -> 0
	var b float32 // Default -> 0
	var c string  // Default -> ""
	var d bool    // Default -> false

	fmt.Println(a, b, c, d)

	// CONSTANTES

	const pi = 3.1416

	fmt.Println(pi)
}
