package main

import (
	base64 "encoding/base64"
	"log"
	"strings"

	resize "github.com/nfnt/resize"
	"image"
	color "image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var (
	// Reference: http://paulbourke.net/dataformats/asciiart/
	asciiMap     = strings.Split("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ", "")
	maxGrayColor = 256
	// Factor to convert 256 color bit grayscale to 70 "color" ascii
	grayscaleToAsciiMultiplier      = 1.0 / float64(maxGrayColor) * float64(len(asciiMap)-1)
	maxImageWidth              uint = 80
)

func Convert(imageBase64 string) (error, string) {
	log.Print("Converting Image")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageBase64))

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Printf("Couldn't decode image %v", err)
		return err, ""
	}

	resizedImage := resizeForAscii(m)
	bounds := resizedImage.Bounds()

	output := ""

	// Convert image to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			colorPixel := resizedImage.At(x, y)
			grayAscii := pixelColorToAscii(colorPixel)
			output = output + grayAscii
		}
		output = output + "\n"
	}
	return nil, output
}

func resizeForAscii(img image.Image) image.Image {

	bounds := img.Bounds()

	newWidth := uint(bounds.Max.X)
	newHeight := uint(bounds.Max.Y / 2)

	if newWidth > maxImageWidth {
		ratio := newWidth / newHeight
		newWidth = maxImageWidth
		newHeight = maxImageWidth / ratio
	}

	return resize.Resize(newWidth, newHeight, img, resize.NearestNeighbor)
}

func pixelColorToAscii(colorPixel color.Color) string {
	grayPixel := color.GrayModel.Convert(colorPixel)
	grayColor := grayPixel.(color.Gray)
	grayAscii := grayToAsciiIndex(grayColor.Y)
	return asciiMap[grayAscii]
}

func grayToAsciiIndex(gray uint8) int {
	asciiFloat := float64(gray) * grayscaleToAsciiMultiplier
	ascii := int(asciiFloat)
	return ascii
}
