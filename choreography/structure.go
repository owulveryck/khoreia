package choreography

import (
	"log"
)

// A Node structure is the base structure of an execution node
type Node struct {
	ID         int               `json:"id",yaml:"id"`
	Name       string            `json:"name",yaml:"name"`
	Target     string            `json:"target",yaml:"target"`
	Interfaces map[string]Action `json:"interfaces",yaml:"interfaces"`
}

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
	var e *FileEngine
	for key, val := range temp {
		log.Println(val)
		engine = key
		e.New(val)
	}
	n.Engine = engine
	n.Interface = e
	return nil
}

func (n *Node) Run() {

}

type Artifact string
type Engine string
type Input string
type Output string

type State int

type Interface interface {
	Do() chan State
	Check() chan bool
	GetOutput() map[string]Output
}

// FileEngine
type FileEngine struct {
	artifact Artifact          `json:"artifact",yaml:"artifact"`
	inputs   []Input           `json:"inputs",yaml:"inputs"`
	outputs  map[string]Output `json:"inputs",yaml:"inputs"`
}

func (e *FileEngine) Check() chan bool {
	c := make(chan bool)
	go func() {
		fileIsPresent := true
		if fileIsPresent {
			c <- true
		} else {
			c <- false
		}
	}()
	return c
}

func (e *FileEngine) New(i map[string]interface{}) {
	for k, v := range i {
		switch k {
		case "artifact":
			e.artifact = Artifact(v.(string))
		}
	}
}

func (e *FileEngine) Do() chan State {
	c := make(chan State)
	log.Printf("Do method of the FileEngine Artifact:%v, Inputs: %v", e.artifact, e.inputs)
	return c
}

func (e *FileEngine) GetOutput() map[string]Output {
	return map[string]Output{}
}
