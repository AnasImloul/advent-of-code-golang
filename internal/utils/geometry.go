package utils

import "github.com/AnasImloul/advent-of-code-golang/internal/types"

func AreAligned[T types.Numeric](point1, point2, point3 types.Point[T]) bool {
	det := point1.X*(point2.Y-point3.Y) +
		point2.X*(point3.Y-point1.Y) +
		point3.X*(point1.Y-point2.Y)
	return det == 0
}

func DistanceSquared[T types.Numeric](point1, point2 types.Point[T]) T {
	dx, dy := point1.X-point2.X, point1.Y-point2.Y
	return dx*dx + dy*dy
}
