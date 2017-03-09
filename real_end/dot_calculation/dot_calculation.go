package dot_calculation

import(
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

func draw_horizontal_line(x1,y,x2 uint32, col color.Color,img *RBGA){
	//it's easy
	for ; x1<=x2;x1++{
		img.Set(x1,y,col)
	}
}

func draw_vertical_line(x,y1,y2 uint32,col color.Color,img *RBGA){
	for ; y1<=y2;y1++{
		img.Set(x,y1,col)
	}
}

func draw_cross(x,y uint32,col color.Color,img *RBGA){
	draw_horizontal_line(x-5,y,x+5,col,img)
	draw_vertical_line(x,y1-5,y1+5)
}

func draw_blue_mark(x,y uint , pic *RBGA,err int){
	offset:= image.Pt(x,y)
	
	draw.Draw(pic,pic.)
}