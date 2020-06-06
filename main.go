package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	mahasiswa "github.com/guaychou/siam-scraper/models"
	"os"
	"strconv"
)
var (
	version string = "v1.0.0"
	url string = "https://siam.ub.ac.id"
	nilaiScale map[string]float32
)

func main() {
	log.Info("Siam scraper version: " + version)
	requestData:= make(map[string]string)

	if os.Getenv("SIAM_NIM")=="" || os.Getenv("SIAM_PASSWORD")=="" {
		cfg,err:=readConfig()
		if err!=nil{
			log.Fatal(err)
		}
		requestData["username"]=cfg.Credentials.Nim
		requestData["password"]=cfg.Credentials.Password
	}
	requestData["username"]=os.Getenv("SIAM_NIM")
	requestData["password"]=os.Getenv("SIAM_PASSWORD")
	requestData["login"]="Masuk"
	scraper:=newScraper()
	nilaiScale=make(map[string]float32)
	nilaiScale["A"]=4.00
	nilaiScale["B+"]=3.50
	nilaiScale["B"]=3.00
	nilaiScale["C+"]=2.50
	nilaiScale["C"]=2.00
	nilaiScale["D+"]=1.50
	nilaiScale["D"]=1.00
	nilaiScale["E"]=0.00
	nilaiScale["K"]=0.00

	err:=loginSiam(scraper,requestData)
	if err!=nil {
		log.Fatal("Login failed\n",err)
	}
	if err!=nil{
		log.Fatal(err)
	}
	data,err:=rekapHasilStudi(scraper)
	if err!=nil {
		log.Fatal(err)
	}
	data.IPK=data.IPK/float32(data.JumlahSKS)
	datajson,err:=json.Marshal(data)
	fmt.Println(string(datajson))
}

func loginSiam(scraper *colly.Collector, data map[string]string) error {
	err:=scraper.Post(url,data)
	if err!=nil{
		return err
	}
	return nil
}


func newScraper() *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36"
	return c
}

func rekapHasilStudi(scraper *colly.Collector) (mahasiswa.DataMahasiswa,error){
	datamahasiswa := mahasiswa.DataMahasiswa{}
	var counter int = 2
	rhs := mahasiswa.RekapHasilStudi{}
	scraper.OnXML("/html/body/table[2]/tbody/tr[1]/td[2]/div[2]", func(e *colly.XMLElement) {
		for {
			if e.ChildText("/table/tbody/tr/td/table/tbody/tr["+strconv.Itoa(counter)+"]")==""{
				break
			}
			rhs.Kode=e.ChildText("/table/tbody/tr/td/table/tbody/tr["+strconv.Itoa(counter)+"]/td[1]")
			rhs.Matkul=e.ChildText("/table/tbody/tr/td/table/tbody/tr["+strconv.Itoa(counter)+"]/td[2]")
			rhs.JumlahSKS=e.ChildText("/table/tbody/tr/td/table/tbody/tr["+strconv.Itoa(counter)+"]/td[3]")
			rhs.Nilai=e.ChildText("/table/tbody/tr/td/table/tbody/tr["+strconv.Itoa(counter)+"]/td[5]")
			datamahasiswa.AddNilai(rhs)
			datamahasiswa.TotalSKS(strconv.Atoi(rhs.JumlahSKS))
			datamahasiswa.HitungIPK(nilaiScale[rhs.Nilai],rhs.JumlahSKS)
			counter++
		}
	})
	scraper.OnHTML(`div[class=text-green]`, func(e *colly.HTMLElement) {
	})
	scraper.OnXML("/html/body/table[2]/tbody/tr[1]/td[2]/table[1]/tbody/tr[1]/td[3]/div/div[1]", func(e *colly.XMLElement) {
		datamahasiswa.Nim=e.Text
	})
	scraper.OnXML("/html/body/table[2]/tbody/tr[1]/td[2]/table[1]/tbody/tr[1]/td[3]/div/div[2]", func(e *colly.XMLElement) {
		datamahasiswa.Nama=e.Text
	})
	scraper.Visit(url+"/rekapstudi.php")
	if datamahasiswa.Nama==""{
		return datamahasiswa,errors.New("Data not found")
	}
	return datamahasiswa,nil
}