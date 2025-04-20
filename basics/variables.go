package basics

import "fmt"

// Variables are mutable

// Global Variables
//middleName := "Cane" // Wrong
var middleName string = "Cane" // accessible to Main Package

func main() {
	//var age int

	//var name string = "John"
	//var name1 = "Jane"

	//count := 10
	//lastName := "Smith"

	middleName = "Mayor"
	fmt.Println(middleName)

	// Default values
	// Numeruc Types: 0
	// Boolean Types: false
	// String Type: ""
	// Pointers, slices, maps, functions, and structs: nil

	// ------ Scope

	//fmt.Println("firstName") // wrong scope

}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}
