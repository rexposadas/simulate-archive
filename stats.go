package simulate

type StatsObj struct {
	Count map[string]int
	MChan chan CountData
}

type CountData struct {
	name  string
	value int
}

// Tick adds a one (+1) to the Map for a given key
func (s *StatsObj) Tick(t string) {

	// todo: implement
	// this function adds (+1) to the map M.
}

func (s *StatsObj) Sub(t string) {
	//todo: implement
}

func (s *StatsObj) Run() {
	s.Count = make(map[string]int)
	s.MChan = make(chan CountData, 1000)
	for {
		m := <-s.MChan

		if val, ok := s.Count[m.name]; ok {
			s.Count[m.name] += val
		} else {
			s.Count[m.name] = 1
		}

	}
	// todo: create a loop which consumes MChan and adds/deletes from Count

}
