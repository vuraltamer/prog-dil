package main
import (
 	"bufio"
  	"fmt"
  	"log"
	"os"
	"time"
	"math/rand"
)

var sozluk = map[string]map[string][]string {
	"en": {
		"name": {
			"Lion",
			"Leopard",
			"Bird",
			"Bear",
			"Monkey",
		},
		"adjective": {
			"Mountain",
			"Grazzy",
			"Crazy",
			"Nervius",
			"Funny",
			"Helpness",
		},
	},
	"tr": {
		"ad": {
			"mimar",
			"adam",
			"ev",
			"bisiklet",
			"bayram",
			"sabah",
			},
		"sifat": {
			"iri",
			"güzel",
			"acayip",
			"fantastik",
                        "kötü",
			"berbat",
		},
	},
}

func rSU(deger int) int { //random sayi uretici
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(deger)
}

func adSecici(dil string) string { // tr ya da en ad seçimi yapar.
	if dil == "tr" {
		return sozluk["tr"]["ad"][rSU(len(sozluk["tr"]["ad"]))]
	}else {
		return sozluk["en"]["name"][rSU(len(sozluk["en"]["name"]))]
	} 
}

func sifatSecici(dil string) string { 
	if dil == "tr" {
		return sozluk["tr"]["sifat"][rSU(len(sozluk["tr"]["sifat"]))]
	}else {
		return sozluk["en"]["adjective"][rSU(len(sozluk["en"]["adjective"]))]
	} 
}

func dilSecici(dil string) string { //random tamlamalar üretir.
	if dil == "tr" {
		return sifatSecici("tr") + " " + adSecici("tr") 
	}else {
		return sifatSecici("en") + " " + adSecici("en")  
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

func readLines(path string) ([]string, error) { //dosyadan veri erişimi...
	file, err := os.Open(path)
	if err != nil {
   		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error { //arrayden dosyaya
	file, err := os.Create(path)
  	if err != nil {
    		return err
  	}
 	defer file.Close()

  	w := bufio.NewWriter(file)
  	for _, line := range lines {
    		fmt.Fprintln(w, line)
  	}
  	return w.Flush()
}

func openfiletoArray(file string) []string { //dosyadan arraya
	lines, err := readLines("aa.txt")
	st := []string{}
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _,line := range lines {
		st = append(st,line)
	}
	return st
}



func tamlama(dil string, file string, sayi int) []string {
	sayac := 0
	cikti := []string{}
	cikti = openfiletoArray(file)
	tam := dilSecici(dil)
	for i := 0; i < 100; i++ {
		if varMi(tam,cikti) == false {
			cikti = append(cikti,tam)
			sayac = sayac + 1
			tam = dilSecici(dil)
		if sayac == sayi {break}
		}
	}
	writeLines(cikti,file)
	return cikti		
}

func main() {
	fmt.Println(tamlama("tr","aa.txt",5))
	//fmt.Println(dilSecici("en"))
}
