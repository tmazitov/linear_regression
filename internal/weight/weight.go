package weight

type Weight struct {
	theta0 float64
	theta1 float64
}

func NewWeight(records [][2]int) *Weight {
	return &Weight{
		theta0: 0.0,
		theta1: 0.0,
	}
}


func (w *Weight) Theta0() float64 { return w.theta0}
func (w *Weight) Theta1() float64 { return w.theta1 }


func (w *Weight) Update(records [][2]int) error {

	var (
		m int = len(records)
	)

	for _, record := range records {
	


	}
	return nil
}
