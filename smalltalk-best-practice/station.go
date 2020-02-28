package smalltalk_best_practice

import "time"

type station struct {
	rate float32
}

type part struct {
	amount float32
	date   time.Time
}

//Bad
func (s *station) ComputePartBad(p part) part {
	return s.multiplyPartTimesRate(p)
}

func (s *station) multiplyPartTimesRate(p part) part {
	return part{amount: p.amount * s.rate, date: p.date}
}

//Better
func (s *station) ComputePartBetter(p part) part {
	return p.multiplyPartTimesRate(s.rate)
}

func (p *part) multiplyPartTimesRate(rate float32) part {
	return part{amount: p.amount * rate, date: p.date}
}
