package main

import (
	"fmt"
	"math"
)

const (
	Speed     = 1.0
	backgroud = " "
	wight     = 160
	height    = 44
	vievScale = 60
	distance  = 100
)

var cubeRenderSlice = make([]string, wight*height)
var cubeDepthSlice = make([]int, wight*height*4)
var A = 0.05
var B = 0.05
var C = 0.05
var Cube = 20.0
var horisontalOffset = 1

func main() {
	//Clear main sreen
	fmt.Println("\x1b[2J")

	for true {
		for i, _ := range cubeRenderSlice {
			cubeRenderSlice[i] = backgroud
		}

		for i, _ := range cubeDepthSlice {
			cubeDepthSlice[i] = 0
		}

		//Render
		for cubeX := -Cube; cubeX < Cube; cubeX += Speed {
			for cubeY := -Cube; cubeY < Cube; cubeY += Speed {
				calcSurface(cubeX, cubeY, -Cube, "@")  //side 1
				calcSurface(Cube, cubeY, cubeX, "$")   //side 2
				calcSurface(-Cube, cubeY, -cubeX, "~") //side 3
				calcSurface(-cubeX, cubeY, Cube, "#")  //side 4
				calcSurface(cubeX, -Cube, -cubeY, ";") //side 5
				calcSurface(cubeX, Cube, cubeY, "+")   //side 6
			}
		}

		fmt.Println("\x1b[H")
		for k := 0; k < wight*height; k++ {
			if k%wight != 0 {
				fmt.Print(cubeRenderSlice[k])
			} else {
				fmt.Println()
			}
		}

		A += 0.05
		B += 0.05
		C += 0.05
	}
}

func calcX(i, j, k float64) float64 {
	return j*math.Sin(A)*math.Sin(B)*math.Cos(C) - k*math.Cos(A)*math.Sin(B)*math.Cos(C) + j*math.Cos(A)*math.Sin(C) + k*math.Sin(A)*math.Sin(C) + i*math.Cos(B)*math.Cos(C)
}

func calcY(i, j, k float64) float64 {
	return j*math.Cos(A)*math.Cos(C) + k*math.Sin(A)*math.Cos(C) - j*math.Sin(A)*math.Sin(B)*math.Sin(C) + k*math.Cos(A)*math.Sin(B)*math.Sin(C) - i*math.Cos(B)*math.Sin(C)
}

func calcZ(i, j, k float64) float64 {
	return k*math.Cos(A)*math.Cos(B) - j*math.Sin(A)*math.Cos(B) + i*math.Sin(B)
}

func calcSurface(cubeX, cubeY, cubeZ float64, ch string) {
	x := calcX(cubeX, cubeY, cubeZ)
	y := calcY(cubeX, cubeY, cubeZ)
	z := calcZ(cubeX, cubeY, cubeZ) + distance

	depth := 1 / z

	matX := int(wight/2 + horisontalOffset + vievScale*depth*x*2)
	matY := int(height/2 + vievScale*depth*y)

	index := matX + matY*wight
	if index >= 0 && index < wight*height {
		if depth > float64(cubeDepthSlice[index]) {
			cubeDepthSlice[index] = int(depth)
			cubeRenderSlice[index] = ch
		}
	}
}
