package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 10 / 2
	restoDaDivisao := 1 % 2

	println(soma)
	println(subtracao)
	println(multiplicacao)
	println(divisao)
	println(restoDaDivisao)

	var numero1 int32 = 10
	var numero2 int32 = 20
	soma2 := numero1 + numero2

	println(soma2)

	var variavel1 string = "ok"
	variavel2 := "ok"
	fmt.Println(variavel1, variavel2)

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	numero := 10
	numero++
	fmt.Println(numero)

	numero += 5
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	numero -= 5
	fmt.Println(numero)

	numero *= 5
	fmt.Println(numero)

	numero /= 5
	fmt.Println(numero)

	numero %= 5
	fmt.Println(numero)

}
