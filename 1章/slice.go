// +build ignore

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

	// n := make([]int, 3, 5)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// n = append(n, 0, 0)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// n = append(n, 1, 2, 3, 4, 5)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	// a := make([]int, 3)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

	// b := make([]int, 0)
	// var c []int
	// fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// var m3 map[string]int
	// m3["pc"] = 5000
	// fmt.Println(m3)

	var s []int
	if s == nil{
		fmt.Println("Nil")
	}
}