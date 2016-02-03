/*
Olivier Wulveryck - author of Gchoreography
Copyright (C) 2015 Olivier Wulveryck

This file is part of the Gchoreography project and
is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this prograv.Digraph.  If not, see <http://www.gnu.org/licenses/>.
*/
package choreography

import (
	"github.com/owulveryck/khoreia/structure"
)

// Graph is the input of the choreography
type Graph struct {
	Name    string           `json:"name,omitempty"`
	State   int              `json:"state"`
	Digraph structure.Matrix `json:"digraph"`
	Nodes   []Node           `json:"nodes"`
	ID      string           `json:"id,omitempty"`
}
