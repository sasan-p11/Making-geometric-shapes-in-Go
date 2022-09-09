package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	f1, err := os.Create("data.html")

	if err != nil {
		log.Fatal(err)
	}

	defer f1.Close()

	var result string

	result = "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='"+strconv.Itoa(width)+"' height='"+strconv.Itoa(height)+"'>"

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			result += "<polygon points='"+strconv.FormatFloat(ax, 'f',-1,32)+","+strconv.FormatFloat(ay, 'f',-1,32)+" "+strconv.FormatFloat(bx, 'f',-1,32)+","+strconv.FormatFloat(by, 'f',-1,32)+" "+strconv.FormatFloat(cx, 'f',-1,32)+","+strconv.FormatFloat(cy, 'f',-1,32)+" "+strconv.FormatFloat(dx, 'f',-1,32)+","+strconv.FormatFloat(dy, 'f',-1,32)+"'/>\n"
		}
	}

	result += "</svg>"
	fmt.Printf("%s\n",result)
	_, err2 := f1.WriteString(result)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
