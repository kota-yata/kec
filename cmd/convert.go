package convert

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"
)

func GetInfoOfSpecifiedImage(srcImagePath string) [6]string{
	srcImgFile, OpenErr := os.Open(srcImagePath)
	if OpenErr != nil {
		log.Fatal(OpenErr)
		os.Exit(1)
	}
	defer srcImgFile.Close()
	fileInfo, StatErr := srcImgFile.Stat()
	if StatErr != nil {
		log.Fatal(StatErr)
		os.Exit(1)
	}
	config, format, DecodeErr := image.DecodeConfig(srcImgFile)
	if DecodeErr != nil {
		log.Fatal(DecodeErr)
		os.Exit(1)
	}
	fileInfoArray := [6] string{
		strconv.FormatBool(fileInfo.IsDir()),
		fileInfo.Name(),
		strconv.FormatInt(fileInfo.Size(), 10),
		strconv.Itoa(config.Width),
		strconv.Itoa(config.Height),
		format,
	}
	return fileInfoArray
}