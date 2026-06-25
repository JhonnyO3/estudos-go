package main

import "fmt"

func main() {

	var (
		slice = []int{1, 2, 3, 4, 5}
		array = [5]int{1, 2, 3, 4, 5}
	)

	fmt.Println("Tamanho Array:", len(array))
	fmt.Println("Capacidade Array:", cap(array))

	fmt.Println("Tamanho Slice:", len(slice))
	fmt.Println("Capacidade Slice:", cap(slice))

	slice = append(slice, 6)

	fmt.Println("-----------: \n", slice)
	fmt.Println("Tamanho Slice:", len(slice))
	fmt.Println("Capacidade Slice:", cap(slice))

}
