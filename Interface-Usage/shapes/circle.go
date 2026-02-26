package shapes

import ("math"
"fmt")

type Circle struct{
	Radius float64
}

func (c *Circle)Area() float64{
	return math.Pi*c.Radius*c.Radius
}

func (c *Circle)PrintShapeDetails(){
	fmt.Println("The total radius of the circle is ",c.Radius)
}