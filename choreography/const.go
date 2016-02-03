//Olivier Wulveryck - author of khoreia
//Copyright (C) 2015 Olivier Wulveryck
//
//This file is part of the khoreia project and
//is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.

package choreography

var States = map[int]string{
	Initial:     "initial",
	Run:         "run",
	Running:     "running",
	Success:     "success",
	Failure:     "failure",
	NotRunnable: "notrunnable",
	Timeout:     "timeout",
	Canceled:    "canceled",
}

const (
	Initial     = 0
	Run         = 1
	Running     = 2
	Success     = 3
	Failure     = 4
	NotRunnable = 5
	Timeout     = 10
	Canceled    = 11
)
