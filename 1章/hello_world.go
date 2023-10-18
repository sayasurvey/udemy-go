// +build ignore

package main

import "fmt"

// init関数は処理の初めに実行される特別な関数
func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

//関数内で他の関数を使用する
func main() {
	bazz()
	fmt.Println("Hello World!", "TEST TEST")
}