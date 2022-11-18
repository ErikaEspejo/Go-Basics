package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	// Imprime sin saltos de linea
	fmt.Print(hola)
	fmt.Print(mundo)

	// Imprime con saltos de linea
	fmt.Println(hola)
	fmt.Println(mundo)

	// Imprime reemplazando valores de variables -> En este link está la info de las variables https://pkg.go.dev/fmt
	nombre := "Erika"
	edad := 26

	fmt.Printf("Hola %s, tu edad es %d \n", nombre, edad) //nombre reemplaza al primer caracter y edad al segundo
	fmt.Printf("Hola %v, tu edad es %v \n", edad, nombre) //nombre reemplaza al primer caracter y edad al segundo

	// Guarda en una variable
	mensajes := fmt.Sprintf("Hola %v, tu edad es %v", edad, nombre)
	fmt.Println(mensajes)

	// Imprimir que tipo de dato tiene una variable
	fmt.Printf("Nombre: %T \n", nombre)

	// Leer y capturar datos de consola
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre) // Captura los datos de consola y los guarda en la variable nombre ya existente, solo captura una palabra sin espacios
	fmt.Println("El nombre modificado es:", nombre)
}
