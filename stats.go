package simulate

type StatsObj struct {
	Count map[string]int
	MChan chan map[string]int
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
	s.MChan = make(chan map[string]int, 1000)

	// todo: create a loop which consumes MChan and adds/deletes from Count

}
