package main

import "fmt"

func tambah(angka1, angka2 int) int {
    return angka1 + angka2
}

func kurang(angka1, angka2 int) int {
	return angka1 - angka2
}

func kali(angka1, angka2 int) int {
	return angka1 * angka2
}

func bagi(angka1, angka2 int) int {
	return angka1 / angka2
}

//func totalNumber(arr []int) int {
	//total := 0
	//for i := 0; i < len(arr); i++ {
		//return total= total + arr[i]
	//}
//}

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

	number2 := []int{
		1,
		2,
		3,
		5,
	}

	number3 := []int{
		0,
		2,
		3,
		4,
	}
	// number[0]

	sum(number)
	sum(number2)
	sum(number3)
}

func sum(a []int) int {
	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	fmt.Println("Totalnya adalah",total)
	return total
}

// func main() {
// 	var hasilTambah int
// 	var hasilKurang int
// 	var hasilKali int
// 	var hasilBagi int

//     hasilTambah = tambah(9,2)
// 	fmt.Println(hasilTambah)

// 	hasilTambah = tambah(10,5)
// 	fmt.Println(hasilTambah)

// 	hasilTambah = tambah(6,3)
// 	fmt.Println(hasilTambah)
	
// 	hasilKurang = kurang(9,2)
// 	fmt.Println(hasilKurang)

// 	hasilKurang = kurang(10,5)
// 	fmt.Println(hasilKurang)

// 	hasilKurang = kurang(6,3)
// 	fmt.Println(hasilKurang)

// 	hasilKali = kali(9,2)
// 	fmt.Println(hasilKali)

// 	hasilKali = kali(10,5)
// 	fmt.Println(hasilKali)

// 	hasilKali = kali(6,3)
// 	fmt.Println(hasilKali)

// 	hasilBagi = bagi(9,2)
// 	fmt.Println(hasilBagi)

// 	hasilBagi = bagi(10,5)
// 	fmt.Println(hasilBagi)

// 	hasilBagi = bagi(6,3)
// 	fmt.Println(hasilBagi)
// }