package simulate

// StatsObject is in charge of keep track of statistics
type StatsObj struct {

	// this variable keeps a map of strings and it's count.
	Count map[string]int
}

// Tick adds a one (+1) to the Map for a given key
func (s *StatsObj) Add(t string) {

	if val, ok := s.Count[t]; ok {
	   newVal := val + 1;
	   s.Count[t] = newVal 
	}

}

func (s *StatsObj) Sub(t string) {

	if val, ok := s.Count[t]; ok {
	   newVal := val - 1;
	   s.Count[t] = newVal 
	}

}

func (s *StatsObj) Run() {
	s.Count = make(map[string]int)

}