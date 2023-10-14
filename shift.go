// +build ignore

package main

import "fmt"

func main() {
	fmt.Println(1 << 0) // 0001 0001 → 1
	fmt.Println(1 << 1) // 0001 0010 → 2
	fmt.Println(1 << 2) // 0001 0100 → 4
	fmt.Println(1 << 3) // 0001 1000 → 8
}
