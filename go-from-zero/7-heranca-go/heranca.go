package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa // Herança
	curso  string
	escola string
}

func main() {

	p1 := pessoa{
		nome:   "Davi",
		idade:  24,
		altura: 180,
	}
	fmt.Println(p1)

	e1 := estudante{
		pessoa: p1,
		curso:  "Engenharia de Software",
		escola: "Universidade de Go",
	}
	fmt.Println(e1)
}
