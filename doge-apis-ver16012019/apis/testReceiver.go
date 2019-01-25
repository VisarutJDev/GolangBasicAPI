package apis

import "fmt"

type Speed struct {
	Distance, Time float64
}

func (s *Speed) SetParam() {
	s.Distance = 100
	s.Time = 5
}

func (s Speed) CalculateSpeed() float64 {
	return s.Distance / s.Time
}

func TestReceiver() string {
	v := &Speed{}
	v.SetParam()
	speed := v.CalculateSpeed()
	return fmt.Sprintf("%s%f %s%f %s%f m/s.", "Distance = ", v.Distance, "m. Time = ", v.Time, "s. Speed = ", speed)
}
