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

	//Zero Values
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	fmt.Printf("Default Int %d, Default Float %f, Default Bool %t, Default String %s\n", defaultInt, defaultFloat, defaultBool, defaultString)

	//go dont have enums, but we can use constants with iota
	const (
		Jan int = iota + 1
		Feb     //2
		Mar     //3
		Apr     //4
	)

	result, r := add(10, 10)
	fmt.Printf("Result %d\n", result, r)

}

//func whit smallcase is private(not exported), with uppercase is public can be exported

func add(a int, b int) (int, int) {
	return a + b, a * b
}
