package punteros

func Swap(px, py *int) {
	a := *px
	b := *py
	*px = b
	*py = a
}
