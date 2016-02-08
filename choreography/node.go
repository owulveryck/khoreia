// Olivier Wulveryck - author of khoreia
// Copyright (C) 2016 Olivier Wulveryck
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

// A Node structure is the base structure of an execution node
type Node struct {
	ID         int                  `json:"id",yaml:"id"`
	Name       string               `json:"name",yaml:"name"`
	Target     string               `json:"target",yaml:"target"`
	Interfaces map[string]Interface `json:"interfaces",yaml:"interfaces"`
	Deps       []Deps               `json:"deps",yaml:"deps"`
	State      State
}

type Deps interface{}

// Lifecycler interface implements a Startup and a Shutdown sequence
type Lifecycler interface {
	Startup()
	Shutdown()
}

// Shutdown executes all the operations <= Delete
func (n *Node) Shutdown() {
	n.Create()
	n.Configure()
	n.Start()
	n.Stop()
	n.Delete()
}

// Startup executes all the operations <= Start
func (n *Node) Startup() {
	n.Create()
	n.Configure()
	n.Start()
}

func (n *Node) Create()    {}
func (n *Node) Configure() {}
func (n *Node) Start()     {}
func (n *Node) Stop()      {}
func (n *Node) Delete()    {}

func (n *Node) Run() {
	switch n.State {
	case Create:
	case Configure:
	case Start:
	case Stop:
	case Delete:
	}
}
