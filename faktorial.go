package main

	import (
		"fmt"
		"math"
	)

	//hitung faktorial dari n
	func hitungFaktorial(n int) float64 {
		if n < 0 {
			panic("n harus bilangan bulat >= 0")
		}
		hasil := 1.0
		for i := 2; i <= n; i++ {
			hasil *= float64(i)
		}
		return hasil
	}

	//faktorial n dibagi 2 pangkat n, lalu bulatkan ke atas
	func faktorialDibagiDuaPangkatN(n int) int {
		faktorial := hitungFaktorial(n)
		pangkatDua := math.Pow(2, float64(n))
		hasilBagi := faktorial / pangkatDua

		// Pembulatan ke atas
		hasilAkhir := int(hasilBagi)
		if hasilBagi > float64(hasilAkhir) {
			hasilAkhir++
		}

		return hasilAkhir
	}

	func main() {
		n := 5
		fmt.Printf("Hasil pembulatan ke atas dari %d! / 2^%d = %d\n", n, n, faktorialDibagiDuaPangkatN(n))
	}
