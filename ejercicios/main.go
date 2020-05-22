package main

import (
	"fmt"

	"github.com/experiments3/ejercicios/cadenas"
)

func main() {
	var cadena1 = "perro"
	var cadena2 = "flauta"
	var posicion = 1
	var resultado = cadenas.Insertln(cadena1, cadena2, posicion)
	fmt.Printf("El resultado es  %v", resultado)
}
