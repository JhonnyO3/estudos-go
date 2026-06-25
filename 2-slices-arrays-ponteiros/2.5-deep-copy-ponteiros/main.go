package main

import "fmt"

type Pessoa struct {
	Nome  *string
	Idade int
}

func main() {
	var nome = "Jhonatas"

	pessoa1 := Pessoa{
		Nome:  &nome,
		Idade: 22,
	}

	pessoa2 := pessoa1

	fmt.Println("Pessoa 1: ", *pessoa1.Nome)
	fmt.Println("Pessoa 2: ", *pessoa2.Nome)

	fmt.Println("-------------------------------")

	pessoa2.Nome = toPointer("Maria")

	fmt.Println("Pessoa 1: ", *pessoa1.Nome)
	fmt.Println("Pessoa 2: ", *pessoa2.Nome)

	fmt.Println("-------------------------------")

	pessoa3 := Pessoa{
		Nome:  toPointer("Jhonny"),
		Idade: 22,
	}

	pessoa4 := deepCopy(pessoa3)

	pessoa4.Nome = toPointer("Carlos")

	fmt.Println("Pessoa 3: ", *pessoa3.Nome)
	fmt.Println("Pessoa 4: ", *pessoa4.Nome)

	fmt.Println("-------------------------------")

	listaDePessoas1 := []Pessoa{
		pessoa1,
		pessoa2,
		pessoa3,
		pessoa4,
	}

	listaDePessoas2 := deepCopyList(listaDePessoas1)

	fmt.Println("Lista 1: ", listaDePessoas1)
	fmt.Println("Lista 2: ", listaDePessoas2)

}

func toPointer(s string) *string {
	return &s
}

func deepCopy(origem Pessoa) Pessoa {
	var destino Pessoa

	destino.Idade = origem.Idade
	destino.Nome = toPointer(*origem.Nome)

	return destino
}

func deepCopyList(origem []Pessoa) []Pessoa {
	var destino = make([]Pessoa, len(origem))

	for i, pessoa := range origem {
		destino[i] = deepCopy(pessoa)
	}
	return destino
}
