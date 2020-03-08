// Provide a set container that stores strings. Because it is
// backed by the standard map, operations should be fairly quick.
package stringset

// A set of strings. StringSet is not thread-safe.
type StringSet interface {
	// Add values to the set. If all strings were previously
	// not part of the set, this method returns true. Otherwise,
	// Put returns false.
	Put(values ...string) (allnew bool)

	// Return whether all values are part of the set.
	Contains(values ...string) bool

	// Remove values from the set. If all strings were previously
	// part of the set, this method returns true. Otherwise,
	// Remove returns false.
	Remove(values ...string) (allremoved bool)

	// Return a copy of all of the strings in the set.
	Strings() []string

	// Return the number of stored strings in this collection.
	Len() int
}

// Return a new empty set.
func New() StringSet {
	return &mapStringSet{
		storage: make(map[string]uint64),
	}
}

// Return a new empty set intialized with values.
func NewWith(values ...string) StringSet {
	set := New()
	set.Put(values...)
	return set
}

type mapStringSet struct {
	storage   map[string]uint64
	nextEpoch uint64
}

func (s *mapStringSet) Put(values ...string) (allnew bool) {
	// get epoch

	epoch := s.nextEpoch
	s.nextEpoch = s.nextEpoch + 1

	// put all values; if we find a value that was added
	// with a previous call to Put, not all values are
	// actually new

	allnew = true

	for _, value := range values {
		if vepoch := s.putSingle(value, epoch); vepoch != epoch {
			allnew = false
		}
	}

	return allnew
}

func (s *mapStringSet) Contains(values ...string) bool {
	for _, value := range values {
		if !s.containsSingle(value) {
			return false
		}
	}

	return true
}

func (s *mapStringSet) Remove(values ...string) (allremoved bool) {
	// values might contain duplicates; to make sure allremoved
	// is correct we need to get ride of those duplicates first

	unique := NewWith(values...).(*mapStringSet)

	// remove all values; if we find a value that is not part
	// of the set, allremoved will be set to false

	allremoved = true

	for value, _ := range unique.storage {
		if !s.removeSingle(value) {
			allremoved = false
		}
	}

	return allremoved
}

func (s *mapStringSet) Strings() []string {
	ss := make([]string, 0, s.Len())
	for value, _ := range s.storage {
		ss = append(ss, value)
	}
	return ss
}

func (s *mapStringSet) Len() int {
	return len(s.storage)
}

// Put a single string into the set. If it was previously not part
// of the set, insert it with map value set to epoch and return epoch.
// If value is already part of the set, return the epoch it was
// added in.
func (s *mapStringSet) putSingle(value string, epoch uint64) uint64 {
	if vepoch, ok := s.storage[value]; !ok {
		s.storage[value] = epoch
		return epoch
	} else {
		return vepoch
	}
}

// Return whether value is part of the set.
func (s *mapStringSet) containsSingle(value string) bool {
	_, contains := s.storage[value]
	return contains
}

// Remove value from set. Returns true if it was previously
// present and now removed. Returns false if it wasn't present.
func (s *mapStringSet) removeSingle(value string) (removed bool) {
	if _, ok := s.storage[value]; ok {
		delete(s.storage, value)
		return true
	} else {
		return false
	}
}
