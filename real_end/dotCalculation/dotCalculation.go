package dotCalculation

//I take care of no fucking error

import (
	"image"
	"image/color"
)

func drawHorizontalLine(x1, y, x2 int, col color.Color, img *(image.RGBA)) {
	//it's easy
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

func drawVerticalLine(x, y1, y2 int, col color.Color, img *(image.RGBA)) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

func drawCross(x, y int, col color.Color, img *(image.RGBA)) {
	drawHorizontalLine(x-5, y, x+5, col, img)
	drawVerticalLine(x, y-5, y+5, col, img)
}

// DrawBlueMark : It draw a blue mark (0,0,255) at (x,y) in the RBGA image img)
func DrawBlueMark(x, y int, img *(image.RGBA)) {
	blue := color.RGBA{0, 0, 255, 255}
	drawCross(100, 100, blue, img)
}
