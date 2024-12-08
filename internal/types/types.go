package types

type Point[T Numeric] struct {
	X, Y T
}

type Numeric interface {
	int | int32 | int64 | float32 | float64
}
