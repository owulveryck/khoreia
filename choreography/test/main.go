package main

import (
	"fmt"
	"github.com/owulveryck/khoreia/choreography"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	var dat []choreography.Node
	f, err := ioutil.ReadFile("example.yaml")

	if err != nil {
		panic(err)

	}

	err = yaml.Unmarshal(f, &dat)
	if err != nil {
		panic(err)
	}
	for _, n := range dat {
		fmt.Println(n)
		//n.Create().Do()

	}

}
