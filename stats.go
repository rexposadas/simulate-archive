package simulate

// StatsObject is in charge of keep track of statistics
type StatsObj struct {

	// this variable keeps a map of strings and it's count.
	Count map[string]int
}

// Tick adds a one (+1) to the Map for a given key
func (s *StatsObj) Add(t string) {
	s.SimpleMath(t, 1)
}

func (s *StatsObj) Sub(t string) {
	s.SimpleMath(t, -1)
}

// simpleMath adds/substracts an arbitrary amount from a map key
func (s *StatsObj) SimpleMath(t string, c int) {

	if val, ok := s.Count[t]; ok {
		s.Count[t] = val + c
	} else {
		s.Count[t] = c
	}
}

func (s *StatsObj) Run() {
	s.Count = make(map[string]int)
}
