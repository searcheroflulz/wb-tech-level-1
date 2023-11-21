package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type ImageProcessor interface {
	processBlackAndWhite() string
}

type ColorProcessor interface {
	processColor() string
}

type BlackAndWhiteImage struct {
	name string
}

func (bw *BlackAndWhiteImage) processBlackAndWhite() string {
	return fmt.Sprint(bw.name, " is black and white")
}

type ColorImage struct {
	name string
}

func (c *ColorImage) processColor() string {
	return fmt.Sprint(c.name, " is colorful")
}

// ColorImageAdapter адаптер для преобразования цветного в черно-белое
type ColorImageAdapter struct {
	colorImage *ColorImage
}

func (ca *ColorImageAdapter) processBlackAndWhite() string {
	return fmt.Sprint(ca.colorImage.name, " was colorful but now it is black and white")
}

func main() {
	blackAndWhiteImage := BlackAndWhiteImage{"GrandPa's photo"}
	colorImage := ColorImage{"Grandchildren's photo"}

	colorImageAdapter := ColorImageAdapter{&colorImage}

	fmt.Println(blackAndWhiteImage.processBlackAndWhite())
	fmt.Println(colorImageAdapter.processBlackAndWhite())
}
