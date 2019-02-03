package apis

import "fmt"

// Formula for moving in a straight line
type Formular interface {
	GetSpeed() float64
}

type FirstSpeedFormular struct {
	EarlySpeed, Acceleration, Time float64
}

// v = u + at
func (s FirstSpeedFormular) GetSpeed() float64 {
	return s.EarlySpeed + (s.Acceleration * s.Time)
}

type SecondSpeedFormular struct {
	EarlySpeed, Acceleration, Distance float64
}

// v^2 = u^2 + 2as
func (s SecondSpeedFormular) GetSpeed() float64 {
	return (s.EarlySpeed * s.EarlySpeed) + (2 * s.Acceleration * s.Distance)
}

func HandleFormular() string {
	first := FirstSpeedFormular{100, 5, 20}
	second := SecondSpeedFormular{10, 5, 100}
	var str string
	allSpeed := make([]Formular, 2)
	allSpeed[0] = first
	allSpeed[1] = second
	for _, speed := range allSpeed {
		str = fmt.Sprintf("%s %f --- ", str, speed.GetSpeed())
	}
	return str
}
