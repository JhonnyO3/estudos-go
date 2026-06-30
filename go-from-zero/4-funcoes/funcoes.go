package main

import "fmt"

func somar(n1, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		return txt
	}

	resultado := f("Olá mundo")
	fmt.Println(resultado)

	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)

}
