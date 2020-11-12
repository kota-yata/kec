package main

import (
	"./cmd"
	"flag"
	"fmt"
	"log"
	"os"
)

func GetFlagsAsArray() []string {
	var args []string
	flag.Parse()
	if argsTmp := flag.Args(); len(argsTmp) >= 1 {
		args = argsTmp
	} else {
		log.Fatal("1 command line argument is required!!!")
		os.Exit(1)
	}
	return args
}

func main() {
	args := GetFlagsAsArray()
	srcImagePath := args[0]
	fileInfoArray := convert.GetInfoOfSpecifiedImage(srcImagePath)
	if len(fileInfoArray) < 6 {
		log.Fatal("1 command line argument is required!!!")
		os.Exit(1)
	}
	fmt.Println("-----Source File Information-----")
	fmt.Printf("Is the relative path indicates Directory? : %v\n", fileInfoArray[0])
	fmt.Printf("File Name : %v\n", fileInfoArray[1])
	fmt.Printf("File Size : %v\n", fileInfoArray[2])
	fmt.Printf("File Width : %vpx\n", fileInfoArray[3])
	fmt.Printf("File Height : %vpx\n", fileInfoArray[4])
	fmt.Printf("File Format : %v\n", fileInfoArray[5])
	fmt.Println("---------------------------------")
}

// 引数は取れたので画像ファイルかどうか判断して画像ファイルだったらフォーマット変換する
// 画像ファイルか判断する関数とフォーマット変換する関数で分けたほうがいい（これをconvert.goに分けてもよい）
// 現状main関数が関数に分けられる余地しかないので先にそれやってmainファイル内整理したほうが良い
