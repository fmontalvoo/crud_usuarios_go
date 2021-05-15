package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os/exec"
)

var reader *bufio.Reader

// Definicion de la estructura.
type Usuario struct {
	id int
	nombre string
	email string
	edad int
}

var id int

// Definicion del mapa.
var usuarios map[int]Usuario

func crearUsuario(){
	limpiarConsola()
	fmt.Println("**Crear usuario**")
	fmt.Print("Ingresa un nombre:")
	nombre := leerTeclado()
	fmt.Print("Ingresa un email:")
	email := leerTeclado()
	fmt.Print("Ingresa una edad:")
	edad, err := strconv.Atoi(leerTeclado())

	if err != nil {
		panic(err)
	}

	id++
	var usuario Usuario = Usuario{id, nombre, email, edad}

	usuarios[id] = usuario
}

func listarUsuario() {
	limpiarConsola()
	fmt.Println("**Listar usuarios**")
	for id, usuario := range usuarios{
		fmt.Println("\nId:", id, "\nNombre:", usuario.nombre, "\nEmail:", usuario.email, "\nEdad:", usuario.edad)
	}
}

func actualizarUsuario() {
	limpiarConsola()
	fmt.Println("**Actualizar usuario**")

	fmt.Print("Ingresa el Id del usuario:")
	id, err := strconv.Atoi(leerTeclado())

	if err != nil {
		panic(err)
	}

	if _, ok := usuarios[id]; ok {

		fmt.Print("Ingresa un nombre:")
		nombre := leerTeclado()
		fmt.Print("Ingresa un email:")
		email := leerTeclado()
		fmt.Print("Ingresa una edad:")
		edad, err := strconv.Atoi(leerTeclado())

		if err != nil {
			panic(err)
		}

		fmt.Println("Usuario:", usuarios[id].nombre, "Actualizado")
		usuario := Usuario{id, nombre, email, edad}
	
		usuarios[id] = usuario
	} else {
		fmt.Println("No existe un usuario con el id:", id )
	}

}

func eliminarUsuario() {
	limpiarConsola()
	fmt.Println("**Eliminar usuario**")
	fmt.Print("Ingresa el Id del usuario:")
	id, err := strconv.Atoi(leerTeclado())

	if err != nil {
		panic(err)
	}

	if _, ok := usuarios[id]; ok {
		fmt.Println("Usuario:", usuarios[id].nombre, "Eliminado")
		delete(usuarios, id)
	} else {
		fmt.Println("No existe un usuario con el id:", id )
	}
}

func leerTeclado()string {

	// Terminal la lectura hasta que el usuario presiona la tecla Enter.
	if opcion, err := reader.ReadString('\n'); err != nil {
		panic("Algo salio mal :(")
	}else{
		// Elimina los saltos de linea de la opcion.
		return strings.TrimSpace(opcion)
	}

}


// Funcion que ejecuta el comando `clear`.
func limpiarConsola() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	var opcion string

	// Inicializamos el mapa.
	usuarios = make(map[int]Usuario)

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Crear")
		fmt.Println("2. Listar")
		fmt.Println("3. Actualizar")
		fmt.Println("4. Eliminar")
		fmt.Print("Opcion: ")
	
		opcion = leerTeclado()

		if opcion == "quit" || opcion == "q"{
			break
		}
	
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

	fmt.Println("Adios...")

}