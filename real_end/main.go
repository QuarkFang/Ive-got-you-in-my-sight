package main

import(
	"./dotCalculation"
	"image"
	"os"
	"image/draw"
	"image/png"
)

func save(img *(image.RGBA), filename string){
	imgw,_:=os.Create(filename)
	png.Encode(imgw,img)
}

func main(){
	imgb,_:=os.Open("Untitled.png")
	GCtest,_:=png.Decode(imgb)
	defer imgb.Close()
	bound := GCtest.Bounds()
	GCtestRGBA:=image.NewRGBA(bound)
	draw.Draw(GCtestRGBA,bound,GCtest,bound.Min,draw.Src)
	col := GCtestRGBA.At(0,0)
	
	save(GCtestRGBA,"fuck.png")
	//dotCalculation.LeaveColor(col,GCtestRGBA)
	save(GCtestRGBA,"fuck2.png")
	
	px,py:=dotCalculation.CalcuGravityCenter(col,GCtestRGBA)
	dotCalculation.DrawBlueMark(px,py,GCtestRGBA)
	save(GCtestRGBA,"fuck3.png")

}

