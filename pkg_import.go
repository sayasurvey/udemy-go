package main

import (
	"fmt"
	"time"
	"os/user"
)

//関数内で他の関数を使用する
func main() {
	fmt.Println("Hello World!", time.Now())
	fmt.Println(user.Current())
}