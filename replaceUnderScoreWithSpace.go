package main
import (
	"os"
    	"fmt"
     	"strings"
 )

type stringYapi struct {
	arg string
}

func (str *stringYapi) replaceUnderScoreWithSpace(){
	
	var ilk int
	var son int
	var i int

	cikti := str.arg
	c_split := strings.Split(cikti, "") // harf harf split ederek listeliyoruz.

	for ilk = 0; ilk < len(c_split); ilk++ { // ilk harfin solundaki ilk alt tirenin indisi. 
		if string(c_split[ilk]) != "_" {break}
	}
	
	for son = len(c_split)-1; son > 0; son-- {  //son harfin sağındaki ilk alt tirenin indisi
		if string(c_split[son]) != "_" {break}
	}
	
	for i = ilk+1; i < son; i++ {
		if  string(c_split[i]) == "_"{
			c_split[i] = " "
		}
	}
	
	cikti = ""
	for _, v := range c_split {     //tekrardan stringe çevirme
		cikti += string(v)
	}
	
	fmt.Println(cikti)

	//////////////////	
}

func main() {
	arg:= stringYapi{arg: os.Args[1]}
      	arg.replaceUnderScoreWithSpace()
}
