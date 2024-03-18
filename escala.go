package main

import "fmt"

const tempK = 373.2

func main() {

	tempC := tempK - 273

	fmt.Printf("A temperatura de ebulição da água é %g °C ou %g °F", tempC, tempK)

}
