package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./for.go")
	if err != nil {
		log.Fatalln(("Error!"))
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	// err = os.Chdir("1章")
	// if err != nil {
	// 	log.Fatalln("Error")
	// }

	err = os.Chdir("1章"); if err != nil {
		log.Fatalln("Error")
	}
}
