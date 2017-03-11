package main

import(
	"./dotCalculation"
	"image"
	"os"
	"image/draw"
	"image/png"
	"image/color"
	"fmt"
)



func save(img *(image.RGBA), filename string){
	imgw,_:=os.Create(filename)
	err := png.Encode(imgw,img)
	if err!=nil{
		println(err)
	}
}

func main(){
	imgb,_:=os.Open("2.png")
	GCtest,_:=png.Decode(imgb)
	defer imgb.Close()
	bound := GCtest.Bounds()
	GCtestRGBA:=image.NewRGBA(bound)
	draw.Draw(GCtestRGBA,bound,GCtest,bound.Min,draw.Src)
	black:=color.RGBA{0,0,0,255}

	col := color.RGBA{227,46,31,255}
	//col := GCtestRGBA.At(715,805)
	col2:=GCtestRGBA.At(0,0)
	fmt.Printf("the target color is %v\n",col)
	
	fmt.Printf("the target color2 is %v\n",col2)

	save(GCtestRGBA,"fuck.png")
	dotCalculation.LeaveColor(col,GCtestRGBA)
	save(GCtestRGBA,"fuck2.png")
	
	px,py:=dotCalculation.CalcuGravityCenter(black,GCtestRGBA)
	dotCalculation.DrawBlueMark(px,py,GCtestRGBA)
	save(GCtestRGBA,"fuck3.png")

}

