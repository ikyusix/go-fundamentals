package main

import "fmt"

func main() {
	// Seleksi kondisi menggunakan keyword if, else if, & else
	var point = 8

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus. Nilai anda %d\n", point)
	}

	// variable temporary pada if-else
	var pointNd = 8840.0

	if percent := pointNd / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// Seleksi kondisi menggunakan keyword switch - case
	// Switch merupakan seleksi konsidi yang sifatnya fokus pada satu variable, lalu kemudian di cek nilainya
	var pointRd = 6
	switch pointRd {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

	// Pemanfaatan case untuk banyak kondisi
	// Sebuah case dapat menampung banyak kondisi. cara penerapannya yaitu dengan menuliskan nilai pembanding-pembandung variable yang diswitch setelah keyword case dipisah tanda koma
	var pointTh = 6
	switch pointTh {
	case 8:
		fmt.Println("Perfect")
	case 7,6,5,4:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

	// Kurung kurawal pada keyword case & default
	var pointFTh = 6
	switch pointFTh {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default: {
		fmt.Println("Not bad")
		fmt.Println("You can be better")
	}
	}

	// Switch dengan gaya if - else
	var pointSTh = 6
	switch {
	case pointSTh == 8:
		fmt.Println("Perfect")
	case (pointSTh < 8) && (pointSTh > 3):
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
		fmt.Println("You need to learn more")
	}

	// Penggunaan keyword fallthrough dalam switch

}
