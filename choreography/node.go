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
	"github.com/kr/pretty"
	//"github.com/owulveryck/khoreia/choreography/engines"
	//"log"
)

// A Node structure is the base structure of an execution node
type Node struct {
	ID         int                  `json:"id",yaml:"id"`
	Name       string               `json:"name",yaml:"name"`
	Target     string               `json:"target",yaml:"target"`
	Interfaces map[string]Interface `json:"interfaces",yaml:"interfaces"`
	Deps       []int                `json:"deps",yaml:"deps"`
}

type Lifecycle interface {
	Create()
	Configure()
	Start()
	Stop()
	Delete()
}

type Interface struct {
	Do    Implementer `json:"do"`
	Check Implementer `json:"check"`
}

// Run calls Check.Check() (which runs in a goroutine) and wait fo all the conditions
// to be ok to call a Do.Do()
func (i *Interface) Run(conditions ...chan bool) chan struct{} {
	check := i.Check.Check()
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
func (n *Node) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type node struct {
		ID         int                    `json:"id",yaml:"id"`
		Name       string                 `json:"name",yaml:"name"`
		Target     string                 `json:"target",yaml:"target"`
		Interfaces map[string]interface{} `json:"interfaces",yaml:"interfaces"`
	}
	//type implementation map[string]map[string]interface{}
	var temp node
	if err := unmarshal(&temp); err != nil {
		return err
	}
	//pretty.Log("DEBUG", pretty.Formatter(temp.Interfaces))
	//var engine string
	//var e Implementation
	for _, val := range temp.Interfaces {
		pretty.Log("DEBUG", pretty.Formatter(val))
		//engine = key
		//f := func(val map[string]interface{}, f func(interface{}) *engines.FileEngine) Implementation {
		//	return f(val)
		//}
		//e = f(val, engines.NewFileEngine)
		//temp.Interfaces[key] = e
	}
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
