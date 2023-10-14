package main

import (
	"fmt"
	"strings"
) 

func main() {
	fmt.Println("Hello World"[0]) //72
	fmt.Println(string("Hello World"[0])) //H

	var s string = "Hello World"
	//変数の内容を変えない場合
	fmt.Println(strings.Replace(s, "H", "X", 1)) //Xello World

	//sの変数の内容も変えたい場合
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	fmt.Println(strings.Contains(s, "World")) //true

	fmt.Println(`Test
		Test
Test`)
}