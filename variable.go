package main

import "fmt"

//関数内で他の関数を使用する
func main() {
	i := 1
	f64 := 1.2
	s := "test"
	t, f := true, false
	fmt.Println(i, f64, s, t, f)
	fmt.Printf("%+v\n", f64)
}