package main

import (
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/owulveryck/khoreia/choreography"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	var nodes []choreography.Node
	f, err := ioutil.ReadFile("samples/topology.yaml")

	if err != nil {
		panic(err)

	}

	err = yaml.Unmarshal(f, &nodes)
	if err != nil {
		log.Println(err)
	}

	// Temp: for debug purpose
	table := uitable.New()
	table.MaxColWidth = 80
	for _, node := range nodes {
		for i, v := range node.Interfaces {
			table.AddRow(node.Name, i)
			v.Run()
		}
	}

	stop := make(chan struct{})
	<-stop
	fmt.Println(table)

}
