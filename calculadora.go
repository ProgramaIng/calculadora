package main

import "fmt"

func main() {
	resultado := calculadora("suma", 0, 0)
	fmt.Println(resultado)
}

func calculadora(operacion string, a, b int) int {

	var resultado int
	switch operacion {
	case "suma":
		resultado = suma(a, b)

	case "resta":
		resultado = resta(a, b)

	case "multiplicacion":
		resultado = multiplicacion(a, b)

	case "division":
		resultado = division(a, b)

	}
	return resultado
}

func suma(x, y int) int {
	return x + y

}

func resta(x, y int) int {
	return x - y

}

func multiplicacion(x, y int) int {
	total := x * y
	return total
}

func division(x, y int) int {
	total := x / y
	return total

}
