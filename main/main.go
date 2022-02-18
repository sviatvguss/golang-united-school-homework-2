package main

import (
	"fmt"
	square "lecture01-solution"
)

func main() {
	fmt.Println()
	fmt.Println(square.CalcSquare(10.0, square.SidesTriangle))
	fmt.Println(square.CalcSquare(10.0, square.SidesSquare))
	fmt.Println(square.CalcSquare(10.0, square.SidesCircle))
}
