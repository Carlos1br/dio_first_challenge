package main

import "fmt"

const kelvin = 373.0

func main() {
	var celsius = kelvin - 273

	fmt.Printf("A temperatura de ebulição da água em ˚K é: %g e a temperatura de ebulição da água em ˚C é %g.", kelvin, celsius)
}
