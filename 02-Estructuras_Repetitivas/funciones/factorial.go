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
