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

	total := sum(number)
	fmt.Println("Total: ", total)

	rata := average2(total, len(number))
	fmt.Println("Rata-rata: ", rata)

	// rataRata := average(number)
	// fmt.Println("Rata rata: ", rataRata)

	min := nilaiMin(number)
	fmt.Println("Nilai minimal ", min)

	max := nilaiMax(number)
	fmt.Println("Nilai maksimal ", max)

	var slice1 = number[0:8]
	fmt.Println("Potongan pertama ", slice1)

	total1 := sum(slice1)
	fmt.Println("Total Potongan pertama ", total1)

	// rataRata1 := average(slice1)
	// fmt.Println("Rata rata potongan pertama: ", rataRata1)

	rata1 := average2(total1, len(slice1))
	fmt.Println("Rata-rata 3: ", rata1)

	min1 := nilaiMin(slice1)
	fmt.Println("Nilai minimal potongan pertama ", min1)

	max1 := nilaiMax(slice1)
	fmt.Println("Nilai maksimal potongan pertama ", max1)

	var slice2 = number[8:16]
	fmt.Println("Potongan kedua ", slice2)

	total2 := sum(slice2)
	fmt.Println("Total Potongan kedua ", total2)

	rata2 := average2(total2, len(slice2))
	fmt.Println("Rata-rata 2: ", rata2)

	// rataRata2 := average(slice2)
	// fmt.Println("Rata rata potongan kedua: ", rataRata2)

	min2 := nilaiMin(slice2)
	fmt.Println("Nilai minimal potongan kedua ", min2)

	max2 := nilaiMax(slice2)
	fmt.Println("Nilai maksimal potongan kedua ", max2)

	var slice3 = number[16:24]
	fmt.Println("Potongan ketiga ", slice3)

	total3 := sum(slice3)
	fmt.Println("Total Potongan ketiga ", total3)

	rata3 := average2(total3, len(slice3))
	fmt.Println("Rata-rata 3: ", rata3)

	// rataRata3 := average(slice3)
	// fmt.Println("Rata rata potongan ketiga: ", rataRata3)

	min3 := nilaiMin(slice3)
	fmt.Println("Nilai minimal potongan ketiga ", min3)

	max3 := nilaiMax(slice3)
	fmt.Println("Nilai maksimal potongan ketiga ", max3)

	var arrTotal = []int{
		total1, // 408 // 0
		total2, // 364 // 1
		total3, // 388 // 2
	}

	var kumpulanSlices = [][]int{
		slice1,
		slice2,
		slice3,
	}

	fmt.Println("Kumpulan slices", kumpulanSlices[0], "punya total ", arrTotal[0])

	//totalMax := totalMax(arrTotal)
	keyTotalMax := getKeyTotalMax(arrTotal)
	fmt.Println("Kumpulan total tertinggi ", kumpulanSlices[keyTotalMax])
}

func sum(a []int) int {
	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	//fmt.Println("Totalnya adalah", total)
	return total
}

func average2(total int, panjang int) int {
	rata := 0
	rata = total / panjang
	return rata
}

// func average(a []int) int {
// 	rata := 0
// 	rata = sum(a) / len(a)
// 	return rata
// }

func nilaiMin(a []int) int {
	var min int
	min = a[0]
	for i := 0; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
	}
	return min
}

func nilaiMax(a []int) int {
	var max int
	max = a[0]
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}

func totalMax(a []int) int {
	var max int
	max = a[0]
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}

func getKeyTotalMax(a []int) int {
	var max int
	max = 0
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = i
		}
	}
	return max
}
