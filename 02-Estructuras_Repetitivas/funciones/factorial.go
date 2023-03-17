package funciones

func Factoreo(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factoreo(n-1)
}
