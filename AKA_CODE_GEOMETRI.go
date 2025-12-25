package main

import "fmt"

// ================= ITERATIF =================
func deretGeometriIteratif(a, r float64, n int) float64 {
	var Sn float64 = 0
	var suku float64
	var i int
	suku = a

	for i = 1; i <= n; i++ {
		Sn = Sn + suku
		suku = suku * r
	}
	return Sn
}

// ================= REKURSIF =================
func deretGeometriRekursif(a, r float64, n int) float64 {
	if n == 1 {
		return a
	}
	return a + deretGeometriRekursif(a*r, r, n-1)
}

// ================= MENU =================
func menu() {
	fmt.Println("\n==== MENU UTAMA ====")
	fmt.Println("1. Deret Geometri Iteratif")
	fmt.Println("2. Deret Geometri Rekursif")
	fmt.Println("3. Exit")
	fmt.Print("Pilih menu: ")
}

// ================= MAIN =================
func main() {
	var a, r float64
	var n, pilih int

	
	fmt.Println("Test Case 1")
	fmt.Println("Iteratif :", deretGeometriIteratif(2, 2, 1))
	fmt.Println("Rekursif:", deretGeometriRekursif(2, 2, 1))

	fmt.Println("\nTest Case 2")
	fmt.Println("Iteratif :", deretGeometriIteratif(2, 2, 4))
	fmt.Println("Rekursif:", deretGeometriRekursif(2, 2, 4))

	fmt.Println("\nTest Case 3")
	fmt.Println("Iteratif :", deretGeometriIteratif(10, 0.5, 4))
	fmt.Println("Rekursif:", deretGeometriRekursif(10, 0.5, 4))

	fmt.Println("\nTest Case 4")
	fmt.Println("Iteratif :", deretGeometriIteratif(5, 1, 3))
	fmt.Println("Rekursif:", deretGeometriRekursif(5, 1, 3))

	fmt.Println("\nTest Case 5")
	fmt.Println("Iteratif :", deretGeometriIteratif(2, -2, 4))
	fmt.Println("Rekursif:", deretGeometriRekursif(2, -2, 4))

	for {
		menu()
		fmt.Scan(&pilih)

		if pilih == 1 {
			fmt.Print("Suku pertama: ")
			fmt.Scan(&a)
			fmt.Print("Rasio: ")
			fmt.Scan(&r)
			fmt.Print("Jumlah suku: ")
			fmt.Scan(&n)

			hasil := deretGeometriIteratif(a, r, n)
			fmt.Println("Jumlah Deret Geometri (Iteratif):", hasil)

		} else if pilih == 2 {
			fmt.Print("Suku pertama: ")
			fmt.Scan(&a)
			fmt.Print("Rasio: ")
			fmt.Scan(&r)
			fmt.Print("Jumlah suku: ")
			fmt.Scan(&n)

			hasil := deretGeometriRekursif(a, r, n)
			fmt.Println("Jumlah Deret Geometri (Rekursif):", hasil)

		} else if pilih == 3 {
			fmt.Println("Program selesai. Terima kasih 🙏")
			break

		} else {
			fmt.Println("Pilihan tidak valid! Masukkan 1–3.")
		}
	}
}
