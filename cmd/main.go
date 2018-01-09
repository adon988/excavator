package main

import (
	"flag"
	"log"
	"strings"

	"github.com/godcong/excavator"
)

var url = flag.String("url", "http://xh.5156edu.com", "catch the web url")

func main() {

	flag.Parse()

	local := *url
	excavator.SetFix(func(s string) string {
		s = strings.Replace(s, "class=font_14", "", -1)
		s = strings.Replace(s, `""`, `"`, -1)
		return s
	})
	excavator.TransformOn()

	list := excavator.GetRootList(local + "/bs.html")

	//log.Println(list)

	excavator.SetFix(func(s string) string {
		s = strings.Replace(s, " bgcolor=#ffffff ", "", -1)
		s = strings.Replace(s, " class=font_14", "", -1)
		s = strings.Replace(s, " bgcolor='#F4F5F9'  align=center", "", -1)
		return s
	})
	log.Println(list)
	//l := excavator.GetCharList("http://xh.5156edu.com/html/908.html")
	//log.Println(l)
	//file, _ := os.OpenFile("excavator.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	//for _, v := range list {
	//	for key, value := range v {
	//		file.WriteString("部首:" + value)
	//		file.WriteString("\n")
	//		list1 := excavator.GetCharList("http://xh.5156edu.com/" + key)
	//
	//		for _, value1 := range list1 {
	//			for _, value2 := range value1 {
	//				file.WriteString(value2)
	//				file.WriteString("\n")
	//			}
	//
	//		}
	//	}
	//
	//}

}