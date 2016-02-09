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
		for i, v := range node.Interfaces {

			log.Printf("calling Do() on %v %v %v", node.Name, i, v)
			v.Do.Do()
			log.Printf("calling Check() on %v %v %v", node.Name, i, v)
			go func(v choreography.Interface) {
				c := v.Check.Check(nil)
				for e := range c {
					log.Printf("Received Event on %v %v %v: %v", node.Name, i, v, e)
					if !e.IsDone {
						log.Printf("Calling v.Do.Do()", v)
						v.Do.Do()
					}
				}
			}(v)
			//v.Run()
		}
	}

	stop := make(chan struct{})
	<-stop
	//time.Sleep(10 * time.Second)
}
