package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 10000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 100
	fmt.Println(numero2)

	//Alias
	// INT32 = RUNE
	var numero3 rune = 1245
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.35
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.35
	fmt.Println(numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	char := 'C'
	fmt.Println(char)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	var boolean bool
	fmt.Println(boolean)
}
