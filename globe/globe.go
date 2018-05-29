package main

import "github.com/mmcloughlin/globe"

func main() {
	g := globe.New()
	g.DrawGraticule(10.0)
	g.DrawLandBoundaries()
	g.CenterOn(51.45, -1)
	g.SavePNG("test.png", 400)

}
