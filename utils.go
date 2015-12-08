package fs

import "encoding/gob"

func Save(filename string, data interface{}) error {
	f, e := Create(filename)
	if e != nil {
		return e
	}
	defer f.Close()
	return gob.NewEncoder(f).Encode(data)
}

func Load(filename string, data interface{}) (interface{}, error) {
	f, e := Open(filename)
	if e != nil {
		return data, e
	}
	defer f.Close()
	return data, gob.NewDecoder(f).Decode(data)
}
