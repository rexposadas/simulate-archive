package simulate

type StatsObj struct {
	Count map[string]int
	MChan chan CountData
}

type CountData struct {
	name  string
	value int
}

func (s *StatsObj) Tick(t string, c int) {

	d := CountData{
		name:  t,
		value: c,
	}
	s.MChan <- d
}

// Add adds a one (+1) to the Map for a given key
func (s *StatsObj) Add(t string) {
	s.Tick(t, 1)
}

// Add substracts one (-1) to the Map for a given key
func (s *StatsObj) Sub(t string) {
	s.Tick(t, -1)
}

func (s *StatsObj) Run() {
	s.Count = make(map[string]int)
	s.MChan = make(chan CountData, 1000)

	for {
		m := <-s.MChan

		if v, ok := s.Count[m.name]; ok {
			s.Count[m.name] += v
		} else {
			s.Count[m.name] = 1
		}
	}
}
