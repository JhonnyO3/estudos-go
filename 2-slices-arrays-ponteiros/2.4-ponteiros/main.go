package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a
	fmt.Println("Variavel a: ", a)
	fmt.Println("Variavel b: ", b)
	fmt.Println("Valor de b: ", *b)

	p := NovaPessoa("Joao", 25)
	fmt.Println(p.Nome, p.Idade, p.Telefone())
	p.AtualizarIdade(26)
	p.atualizarTelefone("4002-8922")
	fmt.Println(p.Nome, p.Idade, p.Telefone())
}

type Pessoa struct {
	Nome     string
	Idade    int
	telefone *string
}

func NovaPessoa(nome string, idade int) Pessoa {
	return Pessoa{
		Nome:  nome,
		Idade: idade,
	}
}

func (p Pessoa) Telefone() string {
	if p.telefone == nil {
		return ""
	}

	return *p.telefone
}

func (p *Pessoa) AtualizarIdade(idade int) {
	p.Idade = idade
}

func (p *Pessoa) atualizarTelefone(telefone string) {
	p.telefone = &telefone
}
