package simulate

// StatsObject is in charge of keep track of statistics
type StatsObj struct {

	// this variable keeps a map of strings and it's count.
	Count map[string]int
}

// Tick adds a one (+1) to the Map for a given key
func (s *StatsObj) Add(t string) {
<<<<<<< HEAD

	if val, ok := s.Count[t]; ok {
	   newVal := val + 1;
	   s.Count[t] = newVal 
	}

=======

	// todo:
	// this function adds the string t to the "Count" member variable. if
	// t already exists in the map, then simply increate the count.
>>>>>>> upstream/master
}

func (s *StatsObj) Sub(t string) {

<<<<<<< HEAD
	if val, ok := s.Count[t]; ok {
	   newVal := val - 1;
	   s.Count[t] = newVal 
	}
=======
	// todo: similar to the Add() method, but substracts a value from Count.
>>>>>>> upstream/master

}

func (s *StatsObj) Run() {
	s.Count = make(map[string]int)

<<<<<<< HEAD
}
=======
}
>>>>>>> upstream/master
