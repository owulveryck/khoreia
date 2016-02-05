package choreography

import (
	"github.com/owulveryck/khoreia/choreography/engines"
)

// A Node structure is the base structure of an execution node
type Node struct {
	ID         int               `json:"id",yaml:"id"`
	Name       string            `json:"name",yaml:"name"`
	Target     string            `json:"target",yaml:"target"`
	Interfaces map[string]Action `json:"interfaces",yaml:"interfaces"`
}

func (n *Node) Create() Interface {
	return n.Interfaces["create"].Do.Interface
}

// TODO: create all other func from the lifecycle

type Action struct {
	Do    Implementation `json:"do",yaml:"do"`
	Check Implementation `json:"check",yaml:"check"`
}

type Implementation struct {
	Engine    string `json:"engine",yaml:"engine"`
	Interface Interface
}

// We need to specialised the Unmarshaling because of the Interfaces field
func (n *Implementation) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type implementation map[string]map[string]interface{}
	var temp implementation
	if err := unmarshal(&temp); err != nil {
		return err
	}
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
}

type Interface interface {
	Do() chan engines.State
	Check() chan bool
	GetOutput() map[string]engines.Output
}
