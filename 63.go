package main

import (
   "code.google.com/p/go-tour/pic"
   "image"
   //"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
   return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
   //var r Rectangle 
   //r.Min.X = 0
   //r.Max.X = 9
   //r.Min.Y = 0
   //r.Max.Y = 9

   r := image.Rect(0, 0, 10, 10)

   return r
}

func (i Image) At(x, y int) color.Color {
   
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
