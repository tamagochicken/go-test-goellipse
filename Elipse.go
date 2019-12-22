package ellipse

import (
	"math"
)

type Init struct {
	A, B float64
}

func (e *Init) GetEccentricity() float64 {
	return (math.Sqrt(math.Pow(e.A, 2) - math.Pow(e.B, 2))) / a
}
