package shapes

import (
	"fmt"
)

type Rectangle struct {
	Length, Width float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r *Rectangle) PrintShapeDetails() {
	fmt.Println("The Width and Length of the rectangle are ", r.Length, r.Width, " lol not respectively")
}
