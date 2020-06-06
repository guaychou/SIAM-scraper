# SIAM Scraper

The purpose of this tools is to scrape your study results

### How to use

```cassandraql
$ export SIAM_NIM=<your college's id>
$ export SIAM_PASSWORD=<your password>
$ ./siam-scraper
```

### Example data

```cassandraql
{
  "Nama": "lordchou",
  "Nim": "xxxxx,
  "IPK": 4.00,
  "JumlahSKS": 150,
  "NilaiMataKuliah": [
    {
      "Kode": "CCE62161",
      "Matkul": "Administrasi Jaringan",
      "JumlahSKS": "3",
      "Nilai": "A"
    },
    {
      "Kode": "CCE62361",
      "Matkul": "Administrasi Sistem Server",
      "JumlahSKS": "3",
      "Nilai": "A"
    },
    {
      "Kode": "CID62125",
      "Matkul": "Statistika",
      "JumlahSKS": "3",
      "Nilai": "A"
    }
  ]
}
```
