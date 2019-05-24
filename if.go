package main

import (
	"fmt"
)

// If
// func main() {

// 	x := 10

// 	if x > 5 {
// 		fmt.Println("x is big")
// 	}

// 	if x > 100 {
// 		fmt.Println("x is very big")
// 	} else {
// 		fmt.Println("x is not that big")
// 	}
// 	if x>5 && x < 15 {
// 		fmt.Println("x is just right")
// 	}
// }

// switch

// func main() {
// 	x := 2

// 	switch x {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	}
// }

//for

// func main() {

// 	for i := 0; i<3; i++ {
// 		fmt.Println(i)
// 	}
// }

//challenge : fizz buzz divisible by 3 print fizz, if divided by 5 print buzz
// by both print fizz buzz

func main() {

	for x := 1; x < 20; x++ {
		if (x%3 == 0) && (x%5 == 0) {
			fmt.Println("fizz buzz")
		} else if x%3 == 0 {
			fmt.Println("fizz")
		} else if x%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(x)
		}
	}
}
