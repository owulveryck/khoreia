package choreography

import (
	"log"
)

// A Node structure is the base structure of an execution node
type Node struct {
	ID         int                       `json:"id",yaml:"id"`
	Name       string                    `json:"name",yaml:"name"`
	Target     string                    `json:"target",yaml:"target"`
	Interfaces map[string]Implementation `json:"interfaces",yaml:"interfaces"`
}

// We need to specialised the Unmarshaling because of the Interfaces field
func (n *Node) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type node struct {
		ID         int                    `json:"id",yaml:"id"`
		Name       string                 `json:"name",yaml:"name"`
		Target     string                 `json:"target",yaml:"target"`
		Interfaces map[string]interface{} `json:"interfaces",yaml:"interfaces"`
	}
	type implementation map[string]map[string]interface{}
	var temp implementation
	if err := unmarshal(&temp); err != nil {
		return err
	}
	log.Println(temp)
	/*
		var engine string
		var e Interface
		for key, val := range temp {
			engine = key
			f := func(val map[string]interface{}, f func(map[string]interface{}) *engines.FileEngine) *engines.FileEngine {
				return f(val)
			}
			e = f(val, engines.NewFileEngine)
		}
		n.Engine = engine
		n.Interface = e
		return nil
	*/
	return nil
}

func (n *Node) RequiredState() chan State {
	states := make(map[string]chan bool, len(n.Interfaces))
	for a, action := range n.Interfaces {
		states[a] = action.Check()
	}
	return nil
}

// Run is a gorouting that will wait for an event and trigger the Do associated
func (n *Node) Run() {
	state := n.RequiredState()
	for state := range state {
		if intf, ok := n.Interfaces[state.(string)]; ok {
			intf.Do()
		} else {
			log.Println("Error, don't know how to go to", state.(string))
		}
	}
}
