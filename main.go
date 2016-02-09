package main

import (
	//"fmt"
	//"github.com/gosuri/uitable"
	"github.com/owulveryck/khoreia/choreography"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	//"time"
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
	for _, node := range nodes {
		for _, v := range node.Interfaces {

			/*
				v.Do.Do()
				go func(v choreography.Interface) {
					c := v.Check.Check(nil)
					for e := range c {
						if !e.IsDone {
							v.Do.Do()
						}
					}
				}(v)
			*/
			v.Run()
		}
	}

	stop := make(chan struct{})
	<-stop
	//time.Sleep(10 * time.Second)
}
