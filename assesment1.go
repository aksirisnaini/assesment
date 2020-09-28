package main

import "fmt"

func main() {
	var number = [24]int{
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
	fmt.Println("Panjang array: ", len(number))
	fmt.Println("Kapasitas array: ", cap(number))

	var slice1 = number[0:8]
	fmt.Println("Potongan pertama ", slice1)

	var total1 int
	var max1 int
	max1 = slice1[0]
	var min1 int
	min1 = slice1[0]
	for i := 0; i < len(slice1); i++ {
		// fmt.Println(slice1[i])
		total1 = total1 + slice1[i]
		if max1 < slice1[i] {
			max1=slice1[i]
		}
		if min1 > slice1[i]{
			min1=slice1[i]
		}
	}
	fmt.Println("Total kumpulan 1 adalah ",total1)
	fmt.Println("Rata-rata kumpulan 1 adalah ",total1/len(slice1))
	fmt.Println("Nilai maksimal kumpulan 1 adalah ", max1)
	fmt.Println("Nilai minimal kumpulan 1 adalah ", min1)


	var slice2 = number[8:16]
	fmt.Println("Potongan kedua ", slice2)

	var total2 int
	var max2 int
	max2 = slice2[0]
	var min2 int
	min2 = slice2[0]
	for i := 0; i < len(slice2); i++ {
		//fmt.Println(slice2[i])
		total2 = total2 + slice2[i]
		if max2 < slice2[i] {
			max2=slice2[i]
		}
		if min2 > slice2[i]{
			min2=slice2[i]
		}
	}
	fmt.Println("Total Kumpulan 2 adalah ",total2)
	fmt.Println("Rata-rata kumpulan 2 adalah ",total2/len(slice2))
	fmt.Println("Nilai maksimal kumpulan 2 adalah ", max2)
	fmt.Println("Nilai minimal kumpulan 2 adalah ", min2)

	var slice3 = number[16:24]
	fmt.Println("Potongan ketiga ", slice3)

	var total3 int
	var max3 int
	max3 = slice3[0]
	var min3 int
	min3 = slice3[0]
	for i := 0; i < len(slice3); i++{
		//fmt.Println(slice3[i])
		total3 = total3 + slice3[i]
		if max3 < slice3[i] {
			max3=slice3[i]
		}
		if min3 > slice3[i]{
			min3=slice3[i]
		}
	}
	fmt.Println("Total Kumpulan 3 ", total3)
	fmt.Println("Rata-rata kumpulan 3 adalah ",total3/len(slice3))
	fmt.Println("Nilai maksimal kumpulan 3 adalah ", max3)
	fmt.Println("Nilai minimal kumpulan 3 adalah ", min3)

	var arrTotal =[] int{
		total1,
		total2,
		total3,
	}
	var max int
	max = arrTotal[0]
	var min int
	min = arrTotal[0]
	for i := 0; i < len(arrTotal); i++ {
		if max < arrTotal[i]{
			max = arrTotal[i]
		}
		if min > arrTotal[i]{
			min = arrTotal[i]
		}
	}
	fmt.Println("Kumpulan total tertinggi ", max)
	fmt.Println("Kumpulan total terendah ", min)
}