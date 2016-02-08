// Olivier Wulveryck - author of khoreia
// Copyright (C) 2015 Olivier Wulveryck
//
// This file is part of the khoreia project and
// is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package choreography

import (
	"fmt"
	"github.com/owulveryck/khoreia/choreography/engines"
	"log"
)

// A Node structure is the base structure of an execution node
type Node struct {
	ID         int                  `json:"id",yaml:"id"`
	Name       string               `json:"name",yaml:"name"`
	Target     string               `json:"target",yaml:"target"`
	Interfaces map[string]Interface `json:"interfaces",yaml:"interfaces"`
	Deps       []Deps               `json:"deps",yaml:"deps"`
}

type Deps interface{}

type Interface struct {
	Do    Implementer `json:"do"`
	Check Implementer `json:"check"`
}

// Run calls Check.Check() (which runs in a goroutine) and wait fo all the conditions
// to be ok to call a Do.Do()
func (i *Interface) Run(conditions ...chan bool) chan struct{} {
	done := make(chan struct{})
	check := i.Check.Check(done)
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				return
			case flag := <-check:
				if flag {
					i.Do.Do()
				}
			}
		}
	}()
	return stop
}

// We need to specialised the Unmarshaling because of the Interfaces field
func (i *Interface) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// attributes is a key/value pair of attributes of an engine
	// eg: filePath = "/tmp/truc"
	type attributes map[string]interface{}
	// the intf is a map of map.
	// the first map's key is Do or Check
	// The second map's key is the engine type
	// The interface{} is all the definition of the engine to be passed to the
	// New function
	type intf map[string]map[string]attributes
	var t intf
	if err := unmarshal(&t); err != nil {
		return err
	}

	// This function takes a map as input
	// The key is the engine name and the input represent the arguments
	f := func(i map[string]attributes) (Implementer, error) {
		var implementer Implementer
		// Get the value of the engine
		for engine, value := range i {
			switch engine {
			case "file":
				impl, err := engines.NewFileEngine(value)
				if err != nil {
					return nil, err
				}
				implementer = impl
			}
		}
		// Create a new implementer based on the engine name

		return implementer, nil
	}
	// key = check | do, otherwise error
	for key, value := range t {
		switch key {
		case "do":
			val, err := f(value)
			if err != nil {
				return err
			}
			i.Do = val
		case "check":
			val, err := f(value)
			if err != nil {
				return err
			}
			i.Check = val
		default:
			return fmt.Errorf("Unknown method %v", key)
		}

	}
	log.Println("DEBUG", t)
	return nil
}

/*
func (n *Node) RequiredState() chan State {
	states := make(map[string]chan bool, len(n.Interfaces))
	for a, action := range n.Interfaces {
		states[a] = action.Check()
	}
	return nil
}

// Run is a gorouting that will wait for an event and trigger the Do associated
func (n *Node) Run() {
	state := n.RequiredState()
	for state := range state {
		if intf, ok := n.Interfaces[state.(string)]; ok {
			intf.Do()
		} else {
			log.Println("Error, don't know how to go to", state.(string))
		}
	}
}
*/
