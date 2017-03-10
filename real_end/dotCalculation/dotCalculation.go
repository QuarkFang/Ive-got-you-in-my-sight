package dotCalculation

//I take care of no fucking error

import (
	"image"
	"image/color"
)

func getDxDy(img *image.RGBA) (int, int) {
	Dx := img.Bounds().Dx()
	Dy := img.Bounds().Dy()
	return Dx, Dy
}

// LeaveColor : all the col in the img will be repalced with white
// then all the other color will be repalced by black, it is for test only
// usage : for test only, to test whther it is possible to use the method
func LeaveColor(col color.Color, img *image.RGBA) {
	Dx, Dy := getDxDy(img)
	for x := 0; x < Dx; x++ {
		for y := 0; y < Dy; y++ {
			if img.At(x, y) != col {
				img.SetRGBA(x, y, color.RGBA{0, 0, 0, 255})
			} else {
				img.SetRGBA(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}
}

// CalcuGravityCenter : calcuGravityCenter of col in the image
func CalcuGravityCenter(col color.Color, img *image.RGBA) (int, int) {
	Dx, Dy := getDxDy(img)
	var sumX, sumY, sumNum int = 0, 0, 0
	for x := 0; x < Dx; x++ {
		for y := 0; y < Dy; y++ {
			if img.At(x, y) == col {
				sumX += x
				sumY += y
				sumNum++
			}
		}
	}
	println(sumX / sumNum)
	println(sumY / sumNum)

	return sumX / sumNum, sumY / sumNum
}

func drawHorizontalLine(x1, y, x2 int, col color.Color, img *image.RGBA) {
	//it's easy
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

func drawVerticalLine(x, y1, y2 int, col color.Color, img *image.RGBA) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

func drawCross(x, y int, col color.Color, img *image.RGBA) {
	drawHorizontalLine(x-10, y, x+10, col, img)
	drawHorizontalLine(x-10, y+1, x+10, col, img)
	drawHorizontalLine(x-10, y-1, x+10, col, img)

	drawVerticalLine(x, y-10, y+10, col, img)

	drawVerticalLine(x-1, y-10, y+10, col, img)
	drawVerticalLine(x+1, y-10, y+10, col, img)
}

// DrawBlueMark : It draw a blue mark (0,0,255) at (x,y) in the RBGA image img)
func DrawBlueMark(x, y int, img *(image.RGBA)) {
	blue := color.RGBA{0, 0, 0, 255}
	drawCross(x, y, blue, img)
}
