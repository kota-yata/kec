package main

import (
	"flag"
	"fmt"
	"image"
		_ "image/png"
	"os"
)

func GetFlagsAsArray() []string{
	var args []string
	flag.Parse()
	if argsTmp := flag.Args(); len(argsTmp) >= 1 {
		args = argsTmp
	}else {
		fmt.Fprintln(os.Stderr, "At least 1 argument is required")
	}
	return args
}

func main() {
	args := GetFlagsAsArray()
	srcImagePath := args[0]
	srcImgFile, err := os.Open(srcImagePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer srcImgFile.Close()
	fileInfo, StatErr := srcImgFile.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, StatErr)
	}
	config, format, DecodeErr := image.DecodeConfig(srcImgFile)
	if DecodeErr != nil {
		fmt.Fprintln(os.Stderr, DecodeErr)
	}
	fmt.Println("-----Source File Information-----")
	fmt.Printf("Is the relative path indicates Directory? : %v\n", fileInfo.IsDir())
	fmt.Printf("File Name : %v\n", fileInfo.Name())
	fmt.Printf("File Size : %v\n", fileInfo.Size())
	fmt.Printf("File Width : %vpx\n", config.Width)
	fmt.Printf("File Height : %vpx\n", config.Height)
	fmt.Printf("File Format : %v\n", format)
	fmt.Println("---------------------------------")
}
