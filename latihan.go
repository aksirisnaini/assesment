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
	// number[0]
	number := add(1, 2)
	number2 := substract(10, 5)
	number3 := pangkatdua(number2)
	number4 := add(pangkatdua(8), pangkatdua(4))
	// 8 dipangkat dua dikurangi hasil dari 10 dikurangi 2
	number5 := substract(pangkatdua(8), substract(10, 2))
	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(number4)
	fmt.Println(number5)
	print(number)
	print(number2)
	
}

func print(a) {
	fmt.Println("Halooo aku mau ngasih tau bahwa aku", a)
}

func add(a int, b int) int {
	result := a + b
	return result
}

func substract( a int, b int) int {
	result := a - b
	return result
}

func pangkatdua(a int) int {
	result := a * a
	return result
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