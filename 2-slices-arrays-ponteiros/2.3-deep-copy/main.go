package main

import "fmt"

func main() {

	fmt.Println("------ Shallow Copy ------")
	matriz1 := [][]int{
		{1, 2},
		{3, 4},
	}

	matriz2 := make([][]int, len(matriz1))
	copy(matriz2, matriz1)

	//matriz2[0] = []int{99, 2}

	//matriz2[0][0] = 99
	fmt.Println(matriz1)
	fmt.Println(matriz2)

	fmt.Println("------ Deep Copy ------")

	matriz3 := [][]int{
		{1, 2},
		{3, 4},
	}

	matriz4 := deepCopy(matriz3)
	matriz4[0][0] = 10

	fmt.Println(matriz3)
	fmt.Println(matriz4)

}

func deepCopy(matriz [][]int) [][]int {
	sliceFormatado := make([][]int, len(matriz))

	for i, valor := range matriz {
		sliceFormatado[i] = make([]int, len(valor))
		copy(sliceFormatado[i], valor)
	}

	return sliceFormatado
}
