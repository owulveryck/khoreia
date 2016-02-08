package choreography

import (
	"fmt"
)

// The nil engine returns true to every check, and do nothing

// NilEngine takes a single file as argument, it checks for its presence or create it
// It relies on the libnotify package
type NilEngine struct {
}

func NewNilEngine(i map[string]interface{}) (*NilEngine, error) {
	for k, v := range i {
		switch k {
		default:
			return nil, fmt.Errorf("Unknown field %v (%v) for  engine NilEngine", k, v)
		}
	}
	return &NilEngine{}, nil
}

// Check if f.File is present and send an event on the channel if it
// appears or disappear
func (f *NilEngine) Check(stop chan struct{}) chan bool {
	c := make(chan bool)

	go func() {
		defer close(c)
		for {
			c <- true
		}
	}()
	return c
}
func (e *NilEngine) Do() {
}

func (e *NilEngine) GetOutput() interface{} {
	return nil
}
