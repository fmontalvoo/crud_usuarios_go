package main

import "fmt"

func crearUsuario(){}

func listarUsuario() {}

func actualizarUsuario() {}

func eliminarUsuario() {}

func main() {

	var opcion string

	fmt.Println("1. Crear")
	fmt.Println("2. Listar")
	fmt.Println("3. Actualizar")
	fmt.Println("4. Eliminar")
	fmt.Println("Opcion: ")
	fmt.Scanf("%s\n", &opcion)

	switch opcion {
		case "1", "crear":
			crearUsuario()
		case "2", "listar":
			listarUsuario()
		case "3", "actualizar":
			actualizarUsuario()
		case "4", "Eliminar":
			eliminarUsuario()
		default:
			fmt.Println("Opcion no valida")
	}
}