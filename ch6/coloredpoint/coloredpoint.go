package coloredpoint

import "image/color"

type Point struct{ X, Y int }

type Coloredpoint struct {
	Point
	Color color.RGBA
}
