package kalkulator

func tambah(a, b float64) float64 {
	return a + b
}

func kurang(a, b float64) float64 {
	return a - b
}

func kali(a, b float64) float64 {
	return a * b
}

func bagi(a, b float64) float64 {
	return a / b
}

func mod(a, b int) int {
	return a % b
}

func Kubus(R float64) (float64, float64) {
	L := 6 * R * R
	V := R * R * R

	return L, V
}

func Balok(p, l, t float64) (float64, float64) {
	L := (2 * p * l) + (2 * l * t) + (2 * p * t)
	V := p * l * t

	return L, V
}

func Tabung(r, t float64) (float64, float64) {
	const phi = 3.14
	L := (2 * phi * r * t) + (2 * phi * r * r)
	V := phi * r * r * t

	return L, V
}