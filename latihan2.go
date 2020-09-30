package main

import "fmt"

func main() {
	fmt.Println("Jalanin program")
	var number = []int{
		23,
		45,
		67,
		54,
		66,
		19,
		56,
		78,
		89,
		44,
		11,
		22,
		33,
		44,
		55,
		66,
		77,
		88,
		99,
		23,
		34,
		32,
		23,
		12,
	}

	fmt.Println(number)
	sum(number)

	var slice1 = number[0:8]
	fmt.Println("Potongan pertama ", slice1)

	sumSlice1(slice1)

	var slice2 = number[8:16]
	fmt.Println("Potongan kedua ", slice2)

	sumSlice2(slice2)

	var slice3 = number[16:24]
	fmt.Println("Potongan ketiga ", slice3)

	sumSlice3(slice3)

}

func sum(a []int) { //tipe data untuk return dihilangkan
	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	fmt.Println("Totalnya adalah", total)
	//return total
}

func sumSlice1(a []int) int {
	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	fmt.Println("Total potongan pertama adalah", total)
	return total
}

func sumSlice2(a []int) int {
	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	fmt.Println("Total potongan kedua adalah", total)
	return total
}

func sumSlice3(a []int) int {
	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	fmt.Println("Total potongan ketiga adalah", total)
	return total
}
