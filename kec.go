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
	if argsTmp := flag.Args(); len(argsTmp) >= 2 {
		args = argsTmp
	} else {
		log.Fatal("2 command line argument is required!!!")
		os.Exit(1)
	}
	return args
}

func main() {
	//CLI description
	KEC := cli.NewApp()
	KEC.Name = "KEC"
	KEC.Usage = "Convert image format to jpg/png/gif"
	KEC.Action = kec
	KEC.Run(os.Args)
}

func kec(c *cli.Context) error {
	args := GetFlagsAsArray()
	srcImagePath := args[0]
	fileInfoArray := convert.GetInfoOfSpecifiedImage(srcImagePath)
	fmt.Println("-----Source File Information-----")
	fmt.Printf("Is the relative path Directory? : %v\n", fileInfoArray[0])
	fmt.Printf("File Name : %v\n", fileInfoArray[1])
	fmt.Printf("File Size : %v\n", fileInfoArray[2])
	fmt.Printf("File Width : %vpx\n", fileInfoArray[3])
	fmt.Printf("File Height : %vpx\n", fileInfoArray[4])
	fmt.Printf("File Format : %v\n", fileInfoArray[5])
	fmt.Println("---------------------------------")

	dstImagePath := args[1]
	widthInt, _ := strconv.Atoi(fileInfoArray[3])
	heightInt, _ := strconv.Atoi(fileInfoArray[4])
	NoEncodedNewImg := convert.SaveNewImgToEncode(srcImagePath, widthInt, heightInt)
	convert.EncodeNewImg(dstImagePath, NoEncodedNewImg)

	return nil
}
