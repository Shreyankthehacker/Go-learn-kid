package main

import(
	"interfaceusage/shapes"
	"fmt"
)

func printShapeArea(s shapes.Shape){
	fmt.Println("The area of the given shape is ",s.Area())
}


func main(){
circle := shapes.Circle{Radius:2.34};
circle.PrintShapeDetails();
printShapeArea(&circle)


}