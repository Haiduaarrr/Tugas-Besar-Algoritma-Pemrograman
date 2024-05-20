package main

import "fmt"

const nmax int = 100

type calonMHS struct {
	nama, domisili, jurusan string
	nilaiTes                int
	status                  string
}

type jurusan struct {
	nama       string
	mahasiswas [15]string // menampung nama mhs
	nMahasiswa int        //
}

type tabJurusan [nmax]jurusan
type tabCalon [nmax]calonMHS

func menu() {
	// Menampilkan pilihan menu
	fmt.Println("===================================")
	fmt.Println("=      PENDAFTARAN MAHASISWA      =")
	fmt.Println("===================================")
	fmt.Println("= 1. Tambah Mahasiswa             =")
	fmt.Println("= 2. Edit Mahasiswa               =")
	fmt.Println("= 3. Delete Mahasiswa             =")
	fmt.Println("= 4. Tampil Data Mahasiswa        =")
	fmt.Println("= 5.                    =")
	fmt.Println("= 6.                    =")
	fmt.Println("Pilih opsi: ")
}

func tambahMhs(mhs *tabCalon, nMhs *int, jur *tabJurusan, nJur int) {
	/*
		IS: Array mhs terdefinisi sembarang
		FS: Array mhs terisi data nama, jurusan, domisili, asal sekolah, dan nilai tes
	*/
	var M calonMHS
	var j jurusan

	fmt.Println("Masukkan data Mahasiswa atau ketik selesai untuk mengakhiri")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&M.nama)
	for M.nama != "selesai" {
		foundJurusan := false
		fmt.Print("Masukkan domisili mahasiswa: ")
		fmt.Scan(&M.domisili)
		fmt.Print("Masukkan jurusan: ")
		fmt.Scan(&M.jurusan)
		for i := 0; i < nJur; i++ {
			if M.jurusan == jur[i].nama {
				jur[i].mahasiswas[nJur] = M.nama
				foundJurusan = true
				j.nMahasiswa++

			}
		}
		if !foundJurusan { // kalo jurusan tidak ditemukan
			fmt.Print("Jurusan tidak ditemukan!")
			M.jurusan = "-"
		}
		fmt.Print("Masukkan nilai tes: ")
		fmt.Scan(&M.nilaiTes)

		mhs[*nMhs] = M
		fmt.Println()
		fmt.Println("Masukkan data Mahasiswa atau ketik selesai untuk mengakhiri")
		fmt.Print("Masukkan nama: ")
		fmt.Scan(&M.nama)
		*nMhs++
	}
}

func cariMahasiswa(mhs tabCalon, n int, nm string) int {
	// Mengembalikan index dari mahasiswa yang dicari
	var idx int
	idx = -1

	for i := 0; i < n; i++ {
		if mhs[i].nama == nm {
			idx = i
		}
	}
	return idx
}

func editMhs(mhs *tabCalon, n *int, nm string) {
	/*
		I.S. Array tabCalon mahasiswa terdefinisi
		F.S. Data nama mahasiswa yang dicari akan terganti
	*/
	var found int
	found = cariMahasiswa(*mhs, *n, nm)
	if found == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Print("Masukan nama mahasiswa: ")
		fmt.Scan(&mhs[found].nama)
		fmt.Print("Masukan domisili mahasiswa: ")
		fmt.Scan(&mhs[found].domisili)
		fmt.Print("Masukan jurusan mahasiswa: ")
		fmt.Scan(&mhs[found].jurusan)
		fmt.Print("Masukan nilai tes mahasiswa: ")
		fmt.Scan(&mhs[found].nilaiTes)
	}
}

func deleteMhs(mhs *tabCalon, n *int, nm string) {
	/*

	 */
	var found int

	found = cariMahasiswa(*mhs, *n, nm)
	if found == -1 {
		fmt.Println("Data belum berhasil di[hapus")
	} else {
		*n = *n - 1
		for i := found; i < *n; i++ {
			mhs[i] = mhs[i+1]
		}
		fmt.Println("Data berhasil dihapus")
	}

}

func tampilMhs(mhs tabCalon, n int) {
	var i int
	fmt.Println("Data Mahasiswa:")
	for i = 0; i < n; i++ {
		fmt.Printf("%s %s %s %d\n", mhs[i].nama, mhs[i].domisili, mhs[i].jurusan, mhs[i].nilaiTes)
	}
	fmt.Println()
}

func tambahJurusan(jur *tabJurusan, nJur *int) {
	var j jurusan

	fmt.Print("Masukkan Data Jurusan atau ketik selesai untuk mengakhiri")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&j.nama)
	for j.nama != "selesai" {
		jur[*nJur] = j
		*nJur++
		fmt.Print("Masukkan nama: ")
		fmt.Scan(&j.nama)
	}
}

func cariJurusan(jur *tabJurusan, nJur int, jurusan string) int {
	// Mengembalikan index dari jurusna yang dicari
	var idx int
	idx = -1

	for i := 0; i < nJur; i++ {
		if jur[i].nama == jurusan {
			idx = i
		}
	}
	return idx
}

func editJurusan(jur *tabJurusan, nJur int, jurusan string) {
	var found int

	found = cariJurusan(jur, nJur, jurusan)
	if found == -1 {
		fmt.Println("Jurusan tidak ditemukan!")
	} else {
		fmt.Print("Masukkan nama jurusan: ")
		fmt.Scan(&jur[found].nama)
		fmt.Println()
		fmt.Print("Data berhasil di edit")
	}

}

func main() {
	//var M calonMHS
	var nMhs, nJur, opsi int
	var mhs tabCalon
	var jur tabJurusan
	var nama string

	menu()
	fmt.Scan(&opsi)
	for opsi != 7 {
		if opsi == 1 {
			tambahMhs(&mhs, &nMhs, &jur, nJur)

		} else if opsi == 2 {
			fmt.Print("Masukkan nama yang datanya ingin diganti: ")
			fmt.Scan(&nama)
			editMhs(&mhs, &nMhs, nama)
		} else if opsi == 3 {
			fmt.Print("Masukkan nama yang datanya ingin dihapus: ")
			fmt.Scan(&nama)
			deleteMhs(&mhs, &nMhs, nama)
		} else if opsi == 4 {
			tampilMhs(mhs, nMhs)
		}
		menu()
		fmt.Scan(&opsi)
	}

}
