package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	convert "./cmd"
	"github.com/urfave/cli"
)

// GetFlagsAsArray Get argument from stdin
func GetFlagsAsArray() []string {
	var args []string
	flag.Parse()
	if argsTmp := flag.Args(); len(argsTmp) >= 3 {
		args = argsTmp
	} else {
		log.Fatal("3 command line argument is required!!!")
		os.Exit(1)
	}
	if args[2] != "compress" && args[2] != "convert" {
		log.Fatal("arg 3 must be 'compress' or 'convert'")
		os.Exit(1)
	}
	return args
}

// ShowImageInfo Show image information
func ShowImageInfo(fileInfoArray [6]string) {
	fmt.Println("-----Source File Information-----")
	fmt.Printf("Is the relative path Directory? : %v\n", fileInfoArray[0])
	fmt.Printf("File Name : %v\n", fileInfoArray[1])
	fmt.Printf("File Size : %v\n", fileInfoArray[2])
	fmt.Printf("File Width : %vpx\n", fileInfoArray[3])
	fmt.Printf("File Height : %vpx\n", fileInfoArray[4])
	fmt.Printf("File Format : %v\n", fileInfoArray[5])
	fmt.Println("---------------------------------")
	fmt.Println("Processing...")
}

func main() {
	//CLI description
	KEC := cli.NewApp()
	KEC.Name = "kec"
	KEC.Usage = "Convert image format to jpg/png/gif"
	KEC.Action = kec
	KEC.Run(os.Args)
}

func kec(c *cli.Context) error {
	args := GetFlagsAsArray()
	srcImagePath := args[0]
	fileInfoArray := convert.GetInfoOfSpecifiedImage(srcImagePath)
	ShowImageInfo(fileInfoArray)
	dstImagePath := args[1]
	compressMode := args[2]
	widthInt, _ := strconv.Atoi(fileInfoArray[3])
	heightInt, _ := strconv.Atoi(fileInfoArray[4])
	NoEncodedNewImg := convert.SaveNewImgToEncode(srcImagePath, widthInt, heightInt, compressMode)
	convert.EncodeNewImg(dstImagePath, NoEncodedNewImg)

	return nil
}
