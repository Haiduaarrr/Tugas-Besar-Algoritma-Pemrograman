package main

import "fmt"

const nmax int = 100

type calonMHS struct {
	nama, domisili, jurusan string
	nilaiTes                int
}

type tabCalon [nmax]calonMHS

func menu() {
	// Menampilkan pilihan menu
	fmt.Println("===================================")
	fmt.Println("=      PENDAFTARAN MAHASISWA      =")
	fmt.Println("===================================")
	fmt.Println("= 1. Tambah Mahasiswa             =")
	fmt.Println("= 2. Edit Mahasiswa               =")
	fmt.Println("= 3.                    =")
	fmt.Println("= 4.                    =")
	fmt.Println("= 5.                    =")
	fmt.Println("= 6.                    =")
	fmt.Println("Pilih opsi: ")
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
	menu()
}

func editMhs(mhs tabCalon, n int, x string) {
	var i int
	var found int = -1
	for found == -1 && i < n {
		if mhs[i].nama == x {
			found = i
		}
		i = i + 1
	}
	if found == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Print("Masukan nama mahasiswa:")
		fmt.Scan(&mhs[i].nama)
		fmt.Print("Masukan domisili mahasiswa:")
		fmt.Scan(&mhs[i].domisili)
		fmt.Print("Masukan jurusan mahasiswa:")
		fmt.Scan(&mhs[i].jurusan)
		fmt.Print("Masukan nilai tes mahasiswa:")
		fmt.Scan(&mhs[i].nilaiTes)
	}
}

func tampilMhs(mhs tabCalon, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("")
	}
}

func main() {
	//var M calonMHS
	var n, opsi int
	var mhs tabCalon

	menu()
	fmt.Scan(&opsi)
	for opsi != 7 {
		if opsi == 1 {
			tambahMhs(&mhs, &n)
		} else if opsi == 2 {

		} else if opsi == 3 {

		}
	}

}
