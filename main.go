package main

import "fmt"

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

func main() {
	LCube, VCube := Kubus(2.5)
	LCuboids, VCuboids := Balok(2,3,5)
	LCylinder, VCylinder := Tabung(7, 5)

	fmt.Println("Hasil Penjumlahan :", tambah(1.5,2))
	fmt.Println("Hasil Pengurangan :", kurang(3,1))
	fmt.Println("Hasil Perkalian :", kali(2,2))
	fmt.Println("Hasil Pembagian :", bagi(3,3))
	fmt.Println("Sisa Pembagian (Modulus) :", mod(3,2))
	fmt.Println("Luas Permukaan Kubus :", LCube)
	fmt.Println("Volume Kubus :", VCube)
	fmt.Println("Luas Permukaan Balok :", LCuboids)
	fmt.Println("Volume Balok :", VCuboids)
	fmt.Println("Luas Permukaan Tabung :", LCylinder)
	fmt.Println("Volume Tabung :", VCylinder)
}