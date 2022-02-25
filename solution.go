package square

import "math"

const (
	SidesTriangle = 3

	SidesSquare = 4

	SidessCircle = 0
)

type intCustomType int

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	if sidesNum == 3 {
		square := (sideLen * sideLen * math.Sqrt(3)) / 4
		return  float64(square)
	} else if sidesNum == 4 {
		square := sideLen * sideLen
		return  float64(square)
	} else if sidesNum == 0 {
		square := sideLen * sideLen * math.Pi
		return  float64(square)
	} else {
		return 0
	}
}
