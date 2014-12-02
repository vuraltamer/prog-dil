package main

 import (
     "fmt"
     "strings"
 )

 func main{
	var str string
	var ilk int
	var son int
	var i int
	fmt.Scanf("%s", str)
	str_split := strings.Split(str, "")

	for ilk = 0; ilk < len(str_split); ilk++ { // ilk harfin solundaki ilk alt tirenin indisi. 
		if string(str_split[ilk]) != "_" {break}
	}
	
	for son = len(str_split)-1; son > 0; son-- {  //son harfin sağındaki ilk alt tirenin indisi
		if string(str_split[son]) != "_" {break}
	}
	
	for i = ilk+1; i < son; i++ {
		if  string(str_split[i]) == "_"{
			str_split[i] = " "
		}
	}
	
	str = ""
	for _, v := range str_split {     //tekrardan stringe çevirme
		str += string(v)
	}
	
	fmt.Println(string(str))

 }
