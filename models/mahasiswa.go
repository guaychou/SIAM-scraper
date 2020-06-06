package mahasiswa

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

type RekapHasilStudi struct {
	Kode string
	Matkul string
	JumlahSKS string
	Nilai string
}

type DataMahasiswa struct {
	Nama string
	Nim string
	IPK float32
	JumlahSKS int
	NilaiMataKuliah []RekapHasilStudi
}

func (mahasiswa *DataMahasiswa) AddNilai (data RekapHasilStudi){
	mahasiswa.NilaiMataKuliah=append(mahasiswa.NilaiMataKuliah,data)
}

func (mahasiswa *DataMahasiswa) HitungIPK (data float32,jumlahSks string){
	intSks,err:=strconv.Atoi(jumlahSks)
	if err!=nil{
		log.Fatal(err)
	}
	mahasiswa.IPK=mahasiswa.IPK+data*float32(intSks)
}

func (mahasiswa *DataMahasiswa) TotalSKS (data int,_ error){
	mahasiswa.JumlahSKS=mahasiswa.JumlahSKS+data
}