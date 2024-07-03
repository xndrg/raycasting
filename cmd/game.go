package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenSize = 1000
	gridSize   = 10
	cellSize   = screenSize / gridSize
)

func main() {
	rl.InitWindow(screenSize, screenSize, "Raycast")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)

	p1 := rl.Vector2{X: 3 * cellSize, Y: 5 * cellSize}
	var mousePosition rl.Vector2
	// var p3 rl.Vector2

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		mousePosition = rl.GetMousePosition()
		// p3 = rayStep(p1, mousePosition)
		rayStep(p1, mousePosition)

		rl.ClearBackground(rl.Black)
		rl.DrawFPS(screenSize-100, 0)
		drawGrid()

		rl.DrawLineEx(p1, mousePosition, 3, rl.LightGray)
		rl.DrawCircleV(p1, 0.2*cellSize, rl.Green)
		rl.DrawCircleV(mousePosition, 0.2*cellSize, rl.Red)
		// rl.DrawCircleLines(int32(p3.X), int32(p3.Y), 0.1*cellSize, rl.Yellow)

		rl.EndDrawing()
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

	d := rl.Vector2Subtract(p2, p1)

	if d.X != 0 {
		k := d.Y / d.X
		c := p1.Y - k*p1.X

		x3v := snap(p2.X, d.X)
		y3v := x3v*k + c
		rl.DrawCircle(int32(x3v*cellSize), int32(y3v*cellSize), 0.1*cellSize, rl.Maroon)

		if k != 0 {
			y3h := snap(p2.Y, d.Y)
			x3h := (y3h - c) / k
			rl.DrawCircle(int32(x3h*cellSize), int32(y3h*cellSize), 0.1*cellSize, rl.DarkBlue)
		}

		return rl.Vector2{X: x3v * cellSize, Y: y3v * cellSize}
	}

	return p2
}

func snap(x float32, dx float32) float32 {
	if dx > 0 {
		return float32(math.Ceil(float64(x)))
	}
	if dx < 0 {
		return float32(math.Floor(float64(x)))
	}

	return x
}
