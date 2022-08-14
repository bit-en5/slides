//START1 OMIT

// package keyvaluestore is a in-memory Key-Value store to save int value according to a string key
package keyvaluestore

import "sync"

// Service that provides the in-memory Key-Value store
type Service struct {
	mux      sync.Mutex
	keyvalue map[string]int
	filename string
}

//END1 OMIT
//START2 OMIT

// New constructs and returns a cache service
func New(dbName string) *Service {
	return &Service{
		keyvalue: make(map[string]int),
		filename: dbName,
	}
}

//END2 OMIT
//START3 OMIT

// Get returns the value related to the given key and a flag to indicate if the key has been found
func (s *Service) Get(key string) (v int, found bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	v, found = s.keyvalue[key]
	return
}

// Set saves the value related to the given key
func (s *Service) Set(key string, value int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.keyvalue[key] = value
}

// Delete removes the value related to the given key
func (s *Service) Delete(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.keyvalue, key)
}

//END3 OMIT
