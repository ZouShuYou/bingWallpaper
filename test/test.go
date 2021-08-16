package main

import (
	"io/ioutil"
	"os"
)

const (
	img2 = "C:/Users/11411/Pictures/Bing/20210816/2.jpg"
	img3 = "C:/Users/11411/Pictures/Bing/20210816/落日时分的香巴拉过山车剪影，西班牙塔拉戈纳萨洛 (© Joaquim F. P./Getty Images).jpg"
)

func main() {
	file, _ := ioutil.ReadFile(img2)
	ioutil.WriteFile(img3,file,os.ModePerm)
}