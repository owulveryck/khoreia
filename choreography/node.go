package choreography

import (
	"github.com/kr/pretty"
	//"github.com/owulveryck/khoreia/choreography/engines"
	//"log"
)

// A Node structure is the base structure of an execution node
type Node struct {
	ID         int                  `json:"id",yaml:"id"`
	Name       string               `json:"name",yaml:"name"`
	Target     string               `json:"target",yaml:"target"`
	Interfaces map[string]Interface `json:"interfaces",yaml:"interfaces"`
	Deps       []int                `json:"deps",yaml:"deps"`
}

type Lifecycle interface {
	Create()
	Configure()
	Start()
	Stop()
	Delete()
}

type Interface struct {
	Do    Implementer `json:"do"`
	Check Implementer `json:"check"`
}

// Run calls Check.Check() (which runs in a goroutine) and wait fo all the conditions
// to be ok to call a Do.Do()
func (i *Interface) Run(conditions ...chan bool) chan struct{} {
	check := i.Check.Check()
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				return
			case flag := <-check:
				if flag {
					i.Do.Do()
				}
			}
		}
	}()
	return stop
}

// We need to specialised the Unmarshaling because of the Interfaces field
func (n *Node) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type node struct {
		ID         int                    `json:"id",yaml:"id"`
		Name       string                 `json:"name",yaml:"name"`
		Target     string                 `json:"target",yaml:"target"`
		Interfaces map[string]interface{} `json:"interfaces",yaml:"interfaces"`
	}
	//type implementation map[string]map[string]interface{}
	var temp node
	if err := unmarshal(&temp); err != nil {
		return err
	}
	//pretty.Log("DEBUG", pretty.Formatter(temp.Interfaces))
	//var engine string
	//var e Implementation
	for _, val := range temp.Interfaces {
		pretty.Log("DEBUG", pretty.Formatter(val))
		//engine = key
		//f := func(val map[string]interface{}, f func(interface{}) *engines.FileEngine) Implementation {
		//	return f(val)
		//}
		//e = f(val, engines.NewFileEngine)
		//temp.Interfaces[key] = e
	}
	return nil
}

/*
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
*/
