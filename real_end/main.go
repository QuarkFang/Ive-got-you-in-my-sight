package main

import(
	"./dotCalculation"
	"image"
	"os"
	"image/jpeg"
)

func save(img *(image.RGBA), filename string){
	imgw,_:=os.Create(filename)
	jpeg.Encode(imgw,img,&jpeg.Options{jpeg.DefaultQuality})
}

func main(){
	screen_size := image.Rect(0,0,200,200)
	test_img :=image.NewRGBA(screen_size)
	
	save(test_img,"fuck.jpeg")
	dotCalculation.DrawBlueMark(200,200,test_img)
	save(test_img,"fuck2.jpeg")
	//save it
}

