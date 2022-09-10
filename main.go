package kalkulator

import "fmt"
 

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