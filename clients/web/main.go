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
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"flag"
	"log"
	"net/http"
)

var OrchestratorUrl string

func main() {

	flag.StringVar(&OrchestratorUrl, "choreography", "http://localhost:8080/v1/tasks", "URL for the choreography")
	flag.Parse()
	router := NewRouter()

	log.Println("connect here: http://localhost:8181/")
	log.Fatal(http.ListenAndServe(":8181", router))
}
