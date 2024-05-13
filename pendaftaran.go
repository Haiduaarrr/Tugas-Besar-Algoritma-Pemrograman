package main

import "fmt"

const nmax int = 100

type calonMHS struct {
	nama, domisili, jurusan, asalSekolah string
	nilaiTes                            int
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

func isicCalon(mhs *tabCalon, *n int) {
	/*
		IS: Array mhs terdefinisi sembarang
  		FS: Array mhs terisi data nama, jurusan, domisili, asal sekolah, dan nilai tes
 	*/
	var M calonMHS

	for M.nama != "selesai" {
		
}
