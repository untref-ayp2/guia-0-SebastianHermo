package calculos

func SumandoNumeros(arreglo []int) (sumatotal int) {
	sumatotal = 0
	for i := 0; i < len(arreglo); i++ {
		sumatotal = sumatotal + arreglo[i]
	}
	return sumatotal
}

func SumaVectorial(arreglo1 []int, arreglo2 []int) int {
	vector1 := 1
	vector2 := 1
	if len(arreglo1) != len(arreglo2) {
		println("Los vectores no son iguales")
	}
	for i := 0; i < len(arreglo1); i++ {
		vector1 = vector1 * arreglo1[i]
	}
	for i := 0; i < len(arreglo2); i++ {
		vector2 = vector2 * arreglo2[i]
	}
	return vector1 + vector2
}
