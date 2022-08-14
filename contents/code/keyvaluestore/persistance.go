package keyvaluestore

import (
	"encoding/json"
	"os"
)

//START1 OMIT

// Save persists the key-value store on disk
func (s *Service) Save() error {
	if s.filename == "" {
		return nil
	}

	data, err := json.Marshal(s.keyvalue)
	if err != nil {
		return err
	}

	return os.WriteFile(s.getFileName(), data, 0777)
}

//END1 OMIT
//START2 OMIT

// Read read the key-value store from disk
func (s *Service) Read() error {
	if s.filename == "" {
		return nil
	}

	data, err := os.ReadFile(s.getFileName())
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.keyvalue)
}

func (s *Service) getFileName() string {
	return os.TempDir() + s.filename
}

//END2 OMIT
