package main

import (
	"estructuras_repetitivas/funciones"
	"fmt"
)

func main() {
	fmt.Println("El Factoreo de 5 es:", funciones.Factoreo(5))
	fmt.Println("El Factoreo de 10 es:", funciones.Factoreo(10))
	fmt.Println("El Factoreo de 0 es:", funciones.Factoreo(0))
	fmt.Println("El producto de 3 y 6 es:", funciones.SumasSucesivas(3, 6))
	fmt.Println("El producto de 0 y 10 es:", funciones.SumasSucesivas(0, 10))
	fmt.Println("El producto de 10 y 100 es:", funciones.SumasSucesivas(10, 100))
	fmt.Println("¿El 9 es primo?", funciones.EsPrimo(9))
	fmt.Println("¿El 11 es primo?", funciones.EsPrimo(11))
	fmt.Println("¿El 1 es primo?", funciones.EsPrimo(1))
}
