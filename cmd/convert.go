package convert

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
)

func GetInfoOfSpecifiedImage(srcImagePath string) [6]string{
	srcImgFile, OpenErr := os.Open(srcImagePath)
	ErrorHandling(OpenErr)
	defer srcImgFile.Close()

	fileInfo, StatErr := srcImgFile.Stat()
	ErrorHandling(StatErr)

	config, format, DecodeErr := image.DecodeConfig(srcImgFile)
	ErrorHandling(DecodeErr)

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