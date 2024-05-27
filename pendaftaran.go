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
	mahasiswas [15]calonMHS // menampung nama mhs
	nMahasiswa int          //
}

type tabJurusan [nmax]jurusan
type tabCalon [nmax]calonMHS

func menuPendaftaran() {
	// Menampilkan pilihan menu
	fmt.Println("===================================")
	fmt.Println("=>     PENDAFTARAN MAHASISWA     <=")
	fmt.Println("===================================")
	fmt.Println("=> 1. Tambah Mahasiswa            =")
	fmt.Println("=> 2. Edit Mahasiswa              =")
	fmt.Println("=> 3. Delete Mahasiswa            =")
	fmt.Println("=> 4. Tampil Data Mahasiswa       =")
	fmt.Println("=> 5. Tambah Jurusan              =")
	fmt.Println("=> 6. Edit Jurusan                =")
	fmt.Println("===================================")
	fmt.Print("Pilih opsi: ")
}

func tambahMhs(mhs *tabCalon, nMhs *int, jur *tabJurusan, nJur int) {
	/*
		I.S.: Array mhs terdefinisi sembarang
		F.S.: Array mhs terisi data nama, jurusan, domisili, asal sekolah, dan nilai tes
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
				jur[i].mahasiswas[nJur] = M
				foundJurusan = true
				j.nMahasiswa++

			}
		}
		if !foundJurusan { // kalo jurusan tidak ditemukan
			fmt.Println("Jurusan tidak ditemukan!")
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

func cariMahasiswa(mhs tabCalon, nMhs int, nm string) int {
	// Mengembalikan index dari nama mahasiswa yang dicari
	var idx int
	idx = -1

	for i := 0; i < nMhs; i++ {
		if mhs[i].nama == nm {
			idx = i
		}
	}
	return idx
}

func editMhs(mhs *tabCalon, nMhs *int, nm string) {
	/*
		I.S. Array tabCalon mahasiswa terdefinisi
		F.S. Data nama mahasiswa yang dicari akan terganti
	*/
	var found int
	found = cariMahasiswa(*mhs, *nMhs, nm)
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

func deleteMhs(mhs *tabCalon, nMhs *int, nm string) {
	/*
		I.S. Array calon mahasiswa mhs dan nMhs terdefinisi
		Proses: Memanggil fungsi cariMahasiswa dan melakukan penimpaan pada idx yang didapat oleh
				fungsi cariMahasiswa
		F.S. Data nama mahasiswa yang dicari telah dihapus
	*/
	var found int

	found = cariMahasiswa(*mhs, *nMhs, nm)
	if found == -1 {
		fmt.Println("Data belum berhasil di[hapus")
	} else {
		*nMhs = *nMhs - 1
		for i := found; i < *nMhs; i++ {
			mhs[i] = mhs[i+1]
		}
		fmt.Println("Data berhasil dihapus")
	}

}

func tampilMhs(mhs tabCalon, nMhs int) {
	/*
		I.S. Array calon mahasiswa mhs dan nMhs terdefinisi
		F.S. Tercetak dilayar data calon mahasiswa
	*/
	var i int
	fmt.Println("Data Mahasiswa:")
	if mhs[i].nama == "" {
		fmt.Println("\nTidak ada data Mahasiswa")
	} else {
		for i = 0; i < nMhs; i++ {
			fmt.Printf("Nama Mahasiswa\t: %s\n", mhs[i].nama)
			fmt.Printf("Domisili Mahasiswa: %s\n", mhs[i].domisili)
			fmt.Printf("Jurusan Mahasiswa: %s\n", mhs[i].jurusan)
			fmt.Printf("Nilai Tes Mahasiswa: %d\n", mhs[i].nilaiTes)
		}
	}

	fmt.Println()
}

func tambahJurusan(jur *tabJurusan, nJur *int) {
	/*
		I.S
		F.S
	*/
	var j jurusan

	fmt.Println("Masukkan Data Jurusan atau ketik selesai untuk mengakhiri")
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
	/*
		I.S
		F.S
	*/
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

func deleteJurusan(jur *tabJurusan, nJur int, jurusan string) {
	/*
		I.S
		F.S
	*/
	var found int

	found = cariJurusan(jur, nJur, jurusan)
	if found == -1 {
		fmt.Println("Jurusan belum berhasil dihapus")
	} else {
		for i := found; i < nJur; i++ {
			jur[i] = jur[i+1]
		}
		fmt.Println("Jurusan berhasil dihapus")
	}
}

func tampilJurusan(jur tabJurusan, nJur int) {
	/*
		I.S
		F.S
	*/
	fmt.Println("Jurusan Yang Tersedia:")
	for i := 0; i < nJur; i++ {
		fmt.Println(jur[i].nama)
	}
	fmt.Println()
}

func sortingNilai(mhs *tabCalon, nMhs int) {
	/*
		I.S Array mhs terdefinisi
		Proses: Mengurutkan nilai mahasiswa dari yang terkecil hingga terbesar (ascending)
		F.S Nilai mahasiswa yang terurut dari yang terkecil hingga terbesar berdasarkan nilai tes
	*/
	var i, j, idx_min int
	var temp calonMHS
	i = 1
	for i <= nMhs-1 {
		idx_min = i - 1
		j = i
		for j < nMhs {
			if mhs[idx_min].nilaiTes > mhs[j].nilaiTes {
				idx_min = j
			}
			j++
		}
		temp = mhs[idx_min]
		mhs[idx_min] = mhs[i-1]
		mhs[i-1] = temp
		i++
	}
}

func sortingNama(mhs *tabCalon, nMhs int) {
	/*
		I.S Array mhs terdefinisi
		Proses: Mengurutkan nama mahasiswa dari yang terkecil hingga terbesar (ascending)
		F.S Nilai mahasiswa yang terurut dari yang terkecil hingga terbesar berdasarkan nama
	*/
	var i, j, idx_min int
	var temp calonMHS
	i = 1
	for i <= nMhs-1 {
		idx_min = i - 1
		j = i
		for j < nMhs {
			if mhs[idx_min].nama > mhs[j].nama {
				idx_min = j
			}
			j++
		}
		temp = mhs[idx_min]
		mhs[idx_min] = mhs[i-1]
		mhs[i-1] = temp
		i++
	}
}

func findMinNilai(mhs tabCalon, nMhs int) int {
	// Mengembalikan nilai minimum dari mahasiswa
	var i, min int
	min = 99999
	for i = 0; i < nMhs; i++ {
		if mhs[i].nilaiTes < min {
			min = mhs[i].nilaiTes
		}
	}
	return min
}

func tampilan() {
	//var M calonMHS
	var nMhs, nJur, opsi1 int
	var mhs tabCalon   // Array untuk mahasiswa
	var jur tabJurusan // Array untuk jurusan
	var namaMahasiswa, namaJurusan string

	menuPendaftaran()
	fmt.Scan(&opsi1)
	for opsi1 != 9 {
		if opsi1 == 1 {
			tambahMhs(&mhs, &nMhs, &jur, nJur)

		} else if opsi1 == 2 {
			fmt.Print("Masukkan nama yang datanya ingin diganti: ")
			fmt.Scan(&namaMahasiswa)
			editMhs(&mhs, &nMhs, namaMahasiswa)
		} else if opsi1 == 3 {
			fmt.Print("Masukkan nama yang datanya ingin dihapus: ")
			fmt.Scan(&namaMahasiswa)
			deleteMhs(&mhs, &nMhs, namaMahasiswa)
		} else if opsi1 == 4 {
			tampilMhs(mhs, nMhs)
		} else if opsi1 == 5 {
			tambahJurusan(&jur, &nJur)
		} else if opsi1 == 6 {
			editJurusan(&jur, nJur, namaJurusan)
		}
		menuPendaftaran()
		fmt.Scan(&opsi1)
	}
}

func main() {
	tampilan()
}
