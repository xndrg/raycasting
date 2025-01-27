package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	raymath "github.com/xndrg/raycast/pkg"
)

var scene [gridSize][gridSize]int = [gridSize][gridSize]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

const (
	screenSize = 1000
	gridSize   = 10
	cellSize   = screenSize / gridSize
	eps        = 1e-3
)

func main() {
	rl.InitWindow(screenSize, screenSize, "Raycast")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)

	p1 := rl.Vector2{X: 3 * cellSize, Y: 5 * cellSize}
	var p2 rl.Vector2
	var p3 rl.Vector2

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		p2 = rl.GetMousePosition()

		rl.ClearBackground(rl.Black)
		rl.DrawFPS(screenSize-100, 0)

		drawWalls()
		drawGrid()

		for {

			p3 = rayStep(p1, p2)
			c := hittingCell(p1, p2)

			p2 = p3

			if c.X/cellSize < 0 || c.X/cellSize >= gridSize ||
				c.Y/cellSize < 0 || c.Y/cellSize >= gridSize ||
				scene[int(c.Y/cellSize)][int(c.X/cellSize)] == 1 {
				break
			}

			rl.DrawCircleLines(int32(p2.X), int32(p2.Y), 0.1*cellSize, rl.Red)
			rl.DrawLineEx(p1, p2, 2, rl.Yellow)

		}

		rl.DrawCircleV(p1, 0.2*cellSize, rl.Green)

		rl.EndDrawing()
	}
}

func drawWalls() {
	var x, y int32

	for y = 0; y < gridSize; y++ {
		for x = 0; x < gridSize; x++ {
			if scene[y][x] != 0 {
				rl.DrawRectangle(
					x*cellSize,
					y*cellSize,
					cellSize,
					cellSize,
					rl.SkyBlue,
				)
			}
		}
	}
}

func drawGrid() {
	for x := 0; x < cellSize; x++ {
		for y := 0; y < cellSize; y++ {
			rl.DrawRectangleLines(
				int32(x*cellSize), int32(y*cellSize),
				cellSize, cellSize,
				rl.DarkBlue,
			)
		}
	}
}

func rayStep(p1 rl.Vector2, p2 rl.Vector2) rl.Vector2 {
	p1 = rl.Vector2{X: p1.X / cellSize, Y: p1.Y / cellSize}
	p2 = rl.Vector2{X: p2.X / cellSize, Y: p2.Y / cellSize}

	p3 := p2
	d := rl.Vector2Subtract(p2, p1)

	if d.X != 0 {
		k := d.Y / d.X
		c := p1.Y - k*p1.X

		x3v := snap(p2.X, d.X)
		y3v := x3v*k + c
		p3 = rl.Vector2{X: x3v, Y: y3v}

		if k != 0 {
			y3h := snap(p2.Y, d.Y)
			x3h := (y3h - c) / k
			p3t := rl.Vector2{X: x3h, Y: y3h}
			if rl.Vector2Distance(p2, p3t) < rl.Vector2Distance(p2, p3) {
				p3 = p3t
			}
		}

	} else {
		y3 := snap(p2.Y, d.Y)
		x3 := p2.X

		p3 = rl.Vector2{X: x3, Y: y3}
	}

	return rl.Vector2{X: p3.X * cellSize, Y: p3.Y * cellSize}
}

func snap(x float32, dx float32) float32 {
	if dx > 0 {
		result := math.Ceil(float64(x) + raymath.Sign(float64(dx))*eps)
		return float32(result)
	}
	if dx < 0 {
		result := math.Floor(float64(x) + raymath.Sign(float64(dx))*eps)
		return float32(result)
	}

	return x
}

func hittingCell(p1, p2 rl.Vector2) rl.Vector2 {
	p1 = rl.Vector2{X: p1.X / cellSize, Y: p1.Y / cellSize}
	p2 = rl.Vector2{X: p2.X / cellSize, Y: p2.Y / cellSize}

	d := rl.Vector2Subtract(p2, p1)
	return rl.Vector2{
		X: float32(math.Floor(float64(p2.X)+raymath.Sign(float64(d.X))*eps) * cellSize),
		Y: float32(math.Floor(float64(p2.Y)+raymath.Sign(float64(d.Y))*eps) * cellSize),
	}
}
