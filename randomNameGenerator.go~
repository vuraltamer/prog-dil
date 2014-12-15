package main
import (
  	"fmt"
	"os"
	"io/ioutil"
	"time"
	"math/rand"
	"strconv"
)

var sozluk = map[string]map[string][]string {
	"en": {
		"ad": {
			"Apricot",
			"April",
			"Ash",
			"Avenue",
			"Ambush",
			"Byte",
			"Ballerina",
			"Banana",
			"Baker",
			"Barrel",
			"Butcher",
			"Course",
			"Cup",
			"Carpenter",
			"Compass",
			"Cell",
			"Diamond",	
			"Dock",
			"Elephant",
			"Axe",
			"File",
			"Lion",
			"Leopard",
			"Bird",
			"Bear",
			"Monkey",
			"Zone",
		},
		"sifat": {
			"Artificial",
			"Attractive",
			"Charming",
			"Boring",
			"Blessed",
			"Convenient",	
			"Famous",
			"Fireless",
			"Poor",
			"Practical",
			"Fine",
			"Mountain",
			"Grazzy",
			"Crazy",
			"Nervius",
			"Funny",
			"Helpness",
			"Banner",
			"Dangerous",
		},
	},
	"tr": {
		"ad": {
			"Adam",
			"Araba",
			"At",
			"Buyruk",
			"Balık",
			"Cam",
			"Çocuk",
			"Dedektif",
			"Durak",
			"Elem",
			"Eylem",
			"Fatura",
			"Halk",
			"Lider",
			"Kalp",
			"Mimar",
			"Adam",
			"Ev",
			"Bisiklet",
			"Buyruk",
			"Sabah",
			},
		"sifat": {
			"Latif",
			"Çirkef",
			"Fedakar",
			"Faydasız",
			"Alakasız",
			"Amansız",
			"Canlı",
			"Heyecanlı",
			"Şirin",
			"Nazlı",
			"Mükemmel",
			"Becerikli",
			"Harika",
			"Harikulade",
			"Gönüllü",
			"Güzel",
			"Acayip",
			"Garip",
                        "Kötü",
			"Berbat",
			"Zor",
			"Zararlı",
		},
	},
}

func rSU(deger int) int { //random sayi uretici
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(deger)
}

func dilTamlamaUreteci(dil string) string { //random tamlamalar üretir.
	if dil == "tr" {
		return sozluk["tr"]["sifat"][rSU(len(sozluk["tr"]["sifat"]))] + " " + sozluk["tr"]["ad"][rSU(len(sozluk["tr"]["ad"]))] 
	}else {
		return sozluk["en"]["sifat"][rSU(len(sozluk["en"]["sifat"]))] + " " + sozluk["en"]["ad"][rSU(len(sozluk["en"]["ad"]))]  
	} 
}

func varMi(ara string, dizin []string) bool { //array içi string arar.
	sonuc := false
	for i := 0; i < len(dizin); i++ {
		if ara == dizin[i] {
			sonuc = true
			break
		}
	}
	return sonuc
}

func createDir(path string, dirName string) {
        os.Mkdir(path + "/" + dirName, 0777)
}

func readDir(path string) []string {
	var list []string
    	files, _ := ioutil.ReadDir(path)
    	for _, f := range files {
		list = append(list, f.Name())
        	
    	}
	return list
}	

func finalFonksiyon(dil string, pathfile string, sayi int) []string {
	sayac := 0
	cikti := []string{}
	eklenen := []string{}
	cikti = readDir(pathfile)
	tam := dilTamlamaUreteci(dil)
	denemesayisi := 0
	if dil == "en" {
		denemesayisi = len(sozluk["en"]["ad"]) * len(sozluk["en"]["sifat"])
	}else {
		denemesayisi = len(sozluk["tr"]["ad"]) * len(sozluk["tr"]["sifat"])
	}
		
	for i := 0; i < denemesayisi; i++ {
		if varMi(tam,cikti) == false {
			cikti = append(cikti, tam)
			eklenen = append(eklenen, tam)
			createDir(pathfile, tam)
			tam = dilTamlamaUreteci(dil)
			sayac = sayac + 1
		if sayac == sayi {break}
		}
	}
	return eklenen		
}

func main() {
	arg:= os.Args[1]
	arg1:= os.Args[2]
	i, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	list := finalFonksiyon(arg,"../"+arg+"_deneme/",i)
	
	for _, n := range list {
		fmt.Println(n)
	}

}
