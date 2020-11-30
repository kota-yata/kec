package convert

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	// underbar
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/image/draw"
)

// OpenImgPath Return file data specified by relative path
func OpenImgPath(srcImagePath string) *os.File {
	srcImgFile, OpenErr := os.Open(srcImagePath)
	ErrorHandling(OpenErr)

	return srcImgFile
}

// GetInfoOfSpecifiedImage Return file information if specified relative path indicates image file
func GetInfoOfSpecifiedImage(srcImagePath string) [6]string {
	srcImgFile := OpenImgPath(srcImagePath)
	defer srcImgFile.Close()
	fileInfo, StatErr := srcImgFile.Stat()
	ErrorHandling(StatErr)

	config, format, DecodeConfigErr := image.DecodeConfig(srcImgFile)
	ErrorHandling(DecodeConfigErr)

	fileInfoArray := [6]string{
		strconv.FormatBool(fileInfo.IsDir()),
		fileInfo.Name(),
		strconv.FormatInt(fileInfo.Size(), 10),
		strconv.Itoa(config.Width),
		strconv.Itoa(config.Height),
		format,
	}

	return fileInfoArray
}

// SaveNewImgToEncode Return redrawn image with Catmull-Rom spline curve
func SaveNewImgToEncode(srcImagePath string, width int, height int) image.Image {
	srcImgFile := OpenImgPath(srcImagePath)
	defer srcImgFile.Close()
	srcImg, _, DecodeErr := image.Decode(srcImgFile)
	ErrorHandling(DecodeErr)

	srcRct := srcImg.Bounds()
	dstImg := image.NewRGBA(image.Rect(0, 0, width, height))
	dstRct := dstImg.Bounds()
	draw.CatmullRom.Scale(dstImg, dstRct, srcImg, srcRct, draw.Over, nil)

	return dstImg
}

// EncodeNewImg Encode new image with specified information
func EncodeNewImg(dstImagePath string, dstImg image.Image) {
	dstImgFile, CreateErr := os.Create(dstImagePath)
	ErrorHandling(CreateErr)
	defer dstImgFile.Close()

	dstFormat := GetFileExtension(dstImagePath)
	switch dstFormat {
	case ".jpg":
		EncodeErr := jpeg.Encode(dstImgFile, dstImg, &jpeg.Options{Quality: 100})
		ErrorHandling(EncodeErr)
	case ".gif":
		EncodeErr := gif.Encode(dstImgFile, dstImg, nil)
		ErrorHandling(EncodeErr)
	case ".png":
		EncodeErr := png.Encode(dstImgFile, dstImg)
		ErrorHandling(EncodeErr)
	}
}

// GetFileExtension Get format extension from end of filename
func GetFileExtension(filePath string) string {
	REGEX := regexp.MustCompile(`(?i)(\.jpg|\.jpeg|\.gif|\.png)$`)
	stringToBeReturned := REGEX.FindString(filePath)
	if stringToBeReturned == "" {
		log.Fatal("2nd argument doesn't have format extension")
		os.Exit(1)
	}
	return stringToBeReturned
}
