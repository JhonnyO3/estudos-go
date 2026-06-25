package main

import "fmt"

func main() {
	fmt.Println("-------------- Copia insegura --------------")

	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = slice1

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	slice1[0] = 10

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	slice2 = append(slice2, 6)
	slice2[0] = 99

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	fmt.Println("-------------- Copia segura --------------")

	var slice3 = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, len(slice3))
	copy(slice4, slice3)

	slice3[0] = 10
	slice4[0] = 25

	fmt.Println("Slice 3:", slice3)
	fmt.Println("Slice 4:", slice4)

}
