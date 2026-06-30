package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
	cidade     string
	estado     string
}

func main() {
	fmt.Println("Arquivo structs")

	endereco1 := endereco{
		logradouro: "Rua dos Devs",
		numero:     10,
		cidade:     "São Paulo",
		estado:     "SP",
	}

	var u usuario
	u.nome = "Davi"
	u.idade = 24
	u.endereco = endereco1
	fmt.Println(u)

	u2 := usuario{
		nome:  "Ana",
		idade: 25,
	}
	fmt.Println(u2)
}
