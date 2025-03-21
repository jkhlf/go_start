package main

//go build -o main main.go, create a binary file named for compiled machine code,, i can specify the plataform with GOOS=linux GOARCH=amd64 go build -o main main.go
//the binary file is like a snapshot of the code, if i change the code i need to compile again
//go run main.go compile and run the code

import (
	"fmt"
)

func main() {

	//Variables
	var name string = "Sung "
	fmt.Printf("My name is %s\n", name)

	//Type inference (Go can infer the type of the variable)
	age := 24
	fmt.Printf("Age %d\n", age)

	var city string
	city = "Japan"
	fmt.Printf("City %s\n", city)

	var country, continent string = "Japan", "Asia"
	fmt.Printf(("Country %s, Continent %s\n"), country, continent)

	var (
		t bool   = true
		s int    = 3000
		p string = "dev"
	)
	fmt.Printf("T %t, S %d, P %s\n", t, s, p)
}
