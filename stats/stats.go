package simulate

// StatsObject is in charge of keep track of statistics
type StatsObj struct {

	// this variable keeps a map of strings and it's count.
	Count map[string]int
}

// Tick adds a one (+1) to the Map for a given key
func (s *StatsObj) Add(t string) {

	s.simpleMath(t, "add")

}

func (s *StatsObj) Sub(t string) {

	s.simpleMath(t, "sub")

}

func (s *StatsObj) simpleMath(t string, o string) {

	newVal := 0
	if val, ok := s.Count[t]; ok {

		// check if operation is add or sub
		if o == "sub" {
			newVal = val - 1
		}else{
			newVal = val + 1
		}   
	}
	s.Count[t] = newVal 
	
}
func (s *StatsObj) Run() {
	s.Count = make(map[string]int)
}