package main

import (
	"fmt"
) 

func main() {
	// n := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(n)
	// fmt.Println(n[2])
	// fmt.Println(n[2:4])
	// fmt.Println(n[:2])
	// fmt.Println(n[2:])
	// fmt.Println(n[:])
	
	// n[2] = 100
	// fmt.Println(n)

	// var board = [][]int{
	// 	{0, 1, 2},
	// 	{3, 4, 5},
	// 	{6, 7, 8},
	// }
	// fmt.Println(board)

	// n = append(n, 100, 200, 300, 400)
	// fmt.Println(n)

	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

	b := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)
}