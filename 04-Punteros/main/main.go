package main

import (
	"fmt"
	"punteros/punteros"
)

func main() {
	x := 5
	y := 7
	px := &x
	py := &y

	fmt.Println("Los Valores de X e Y son:", *px, *py)
	punteros.Swap(px, py)
	fmt.Println("Los valores de X e Y son:", *px, *py)
}
