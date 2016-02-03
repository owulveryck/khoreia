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
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/owulveryck/khoreia/choreography"
	"html/template"
	"net/http"
)

type jsonErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func displayGraph(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string
	id = vars["id"]
	g, err := getGraph(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: fmt.Sprintf("%v", err)}); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(g); err != nil {
		panic(err)
	}
}

func displayMain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string
	id = vars["id"]
	g, err := getGraph(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: fmt.Sprintf("%v", err)}); err != nil {
			panic(err)
		}
		return
	}
	type res struct {
		ID     string
		Update string
		Nodes  []choreography.Node
	}

	t := template.New("template.tmpl")
	t, err = t.ParseFiles("tmpl/template.tmpl", "tmpl/viewgraph.tmpl", "tmpl/navbar_right.tmpl")
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: fmt.Sprintf("%v", err)}); err != nil {
			panic(err)
		}
		return
	}
	//w.WriteHeader(http.StatusOK)
	var nodes []choreography.Node
	for i, _ := range g.Nodes {
		if g.Nodes[i].Artifact != "" || g.Nodes[i].Engine != "nil" {
			nodes = append(nodes, g.Nodes[i])
		}
	}
	err = t.Execute(w, res{id, fmt.Sprintf("/graph/%v.json", id), nodes})
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: fmt.Sprintf("%v", err)}); err != nil {
			panic(err)
		}
		return
	}
}
func getOrchestratorState(w http.ResponseWriter, r *http.Request) {

}
func getTasks(w http.ResponseWriter, r *http.Request) {
}
func displayTasks(w http.ResponseWriter, r *http.Request) {

}

func displaySvg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id string
	id = vars["id"]
	b, err := getSvg(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Msg: "Not Found"}); err != nil {
			panic(err)

		}
	}
	w.Header().Set("Content-Type", "image/svg+xml; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", b)
}
