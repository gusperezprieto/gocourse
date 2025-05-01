package main

import "fmt"

func main() {

	// Switch statement
	// switch expresion {
	// case value1:
	// 	// code
	// case value2:
	// 	//code
	// default:
	// 	//code
	// }

	// Switch statement
	// switch expresion {
	// case value1:
	// 	// code
	//	fallthrough (to evaluate the following case)
	// case value2:
	// 	//code
	// default:
	// 	//code
	// }

	// fruit := "pinapple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple")
	// case "banana":
	// 	fmt.Println("It's a banana")
	// default:
	// 	fmt.Println("Unknow fruit")
	// }

	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")
	// case "Sturday", "Sunday":
	// 	fmt.Println("It's Weekend")
	// default:
	// 	fmt.Println("Invalid Day")
	// }

	// num := 2
	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not 2")
	// }

	checkType(10)
	checkType(2.32)
	checkType("Hola")
	checkType(true)

}

func checkType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("It's integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unkown type")

	}
}
