package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	var dat interface{}
	if err := dec.Decode(&dat); err != nil {
		panic(err)
	}
	d, err := yaml.Unmarshal(&dat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", d)

}
