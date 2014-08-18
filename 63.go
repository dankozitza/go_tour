package main

import (
   "code.google.com/p/go-tour/pic"
   "image"
   "image/color"
)

type Image struct{
   graph [][]color.Color
}

func (i Image) ColorModel() color.Model {
   return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
   //var r Rectangle 
   //r.Min.X = 0
   //r.Max.X = 9
   //r.Min.Y = 0
   //r.Max.Y = 9

   r := image.Rect(0, 0, 255, 255)

   return r
}

func (i Image) At(x, y int) color.Color {
   return i.graph[y][x]
}

func generate_colors(img *Image) {
 
    img.graph := make([][]img.colorModel, img.Bounds.Max.Y)
 
    for y := range img.graph {
       img.graph[y] = make([]img.colorModel, img.Bounds.Max.X)
       for x := range graph[y] {
          img.graph[y][x] = img.colorModel{x, y, 255, 255}
       }
    }
}

func main() {
    m := Image{}
    generate_colors(m)
    pic.ShowImage(m)
}
