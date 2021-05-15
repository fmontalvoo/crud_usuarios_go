package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func crearUsuario(){}

func listarUsuario() {}

func actualizarUsuario() {}

func eliminarUsuario() {}

func main() {

	var opcion string
	var reader *bufio.Reader

	reader = bufio.NewReader(os.Stdin)

	fmt.Println("1. Crear")
	fmt.Println("2. Listar")
	fmt.Println("3. Actualizar")
	fmt.Println("4. Eliminar")
	fmt.Print("Opcion: ")

	// Terminal la lectura hasta que el usuario presiona la tecla Enter.
	opcion, err := reader.ReadString('\n') 

	if err != nil {
		panic("Algo salio mal :(")
	}

	// Elimina los saltos de linea de la opcion.
	opcion = strings.TrimSpace(opcion)

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
			fmt.Println("Opcion:", opcion, "no valida")
	}
}