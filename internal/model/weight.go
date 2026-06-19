package model

type Weight struct {
	theta0 float64
	theta1 float64
}

func NewWeight() *Weight {
	return &Weight{
		theta0: 0.0,
		theta1: 0.0,
	}
}

func (w *Weight) Theta0() float64 { return w.theta0 }
func (w *Weight) Theta1() float64 { return w.theta1 }

func (w *Weight) Update(thetaPair [2]float64) {
	w.theta0 = thetaPair[0]
	w.theta1 = thetaPair[1]
}
