package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// есть принтер, который печатает только черно-белые фото
type BlackAndWhitePrinter struct{}

func (bwp *BlackAndWhitePrinter) PrintBlackAndWhite(img Image) {
	fmt.Println(img.processBlackAndWhite())
}

// черно-белый принтер принимает данный интерфейс
type Image interface {
	processBlackAndWhite() string
}

type BlackAndWhiteImage struct {
	name string
}

func (bw BlackAndWhiteImage) processBlackAndWhite() string {
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

// данной функцией мы удовлетворяем требованиям интерфейса Image
func (ca ColorImageAdapter) processBlackAndWhite() string {
	return fmt.Sprint(ca.colorImage.name, " was colorful but now it is black and white")
}

func main() {
	blackAndWhiteImage := BlackAndWhiteImage{"GrandPa's photo"}
	colorImage := ColorImage{"Grandchildren's photo"}

	colorImageAdapter := ColorImageAdapter{&colorImage}
	printer := BlackAndWhitePrinter{}

	printer.PrintBlackAndWhite(blackAndWhiteImage)
	printer.PrintBlackAndWhite(colorImageAdapter)
}
