package funciones

func Factoreo(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factoreo(n-1)
}
