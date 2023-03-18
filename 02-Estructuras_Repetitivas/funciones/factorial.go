package funciones

func Factoreo(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factoreo(n-1)
}

func SumasSucesivas(a int, b int) int {
	return ((a * (a + 1)) / 2) * ((b * (b + 1)) / 2)
}

func EsPrimo(a int) bool {
	resultado := 1
	NumeroPrimo := false
	numerosPrimos := [5]int{a, 2, 3, 5, 7}

	for i := 0; i < 5; i++ {
		if a%numerosPrimos[i] == 0 {
			resultado = resultado + 1
		}
	}
	if resultado == 2 {
		NumeroPrimo = true
	}
	return NumeroPrimo
}
