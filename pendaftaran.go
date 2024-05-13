package main

import "fmt"

const nmax int = 100

type calonMHS struct {
	nama, domisili, jurusan string
	nilaiTes	int
}

type tabCalon [nmax]calonMHS
 
func menu() {
	// Menampilkan pilihan menu
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func tambahMhs(mhs *tabCalon, n *int) {
	/*
		IS: Array mhs terdefinisi sembarang
  		FS: Array mhs terisi data nama, jurusan, domisili, asal sekolah, dan nilai tes
 	*/
	var M calonMHS

	fmt.Print("Masukkan nama mahasiswa atau ketik selesai untuk mengakhiri: ")
	fmt.Scan(&M.nama)
	for M.nama != "selesai" {
		fmt.Print("Masukkan domisili mahasiswa: ")
		fmt.Scan(&M.domisili)
		fmt.Print("Masukkan jurusan: ")
		fmt.Scan(&M.jurusan)
		fmt.Print("Masukkan nilai tes: ")
		fmt.Scan(&M.nilaiTes)

		mhs[*n] = M
		fmt.Print("Masukkan nama mahasiswa atau ketik selesai untuk mengakhiri: ")
	    fmt.Scan(&M.nama)
		*n++
	}
}

func main() {
    //var M calonMHS
    var n int
    var mhs tabCalon
    
    tambahMhs(&mhs, &n)
    
}

