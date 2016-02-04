package choreography

// A Node structure is the base structure of an execution node
type Node struct {
	ID         int        `json:"id",yaml:"id"`
	Name       string     `json:"name",yaml:"name"`
	Target     string     `json:"target",yaml:"target"`
	Interfaces Interfaces `json:"interfaces",yaml:"interfaces"`
}

// Interfaces is a type used for adding a special unmarshaling func
type Interfaces map[string]Interface

// We need to specialised the Unmarshaling because of the Interfaces field
func (n *Interfaces) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var temp map[string]interface{}
	if err := unmarshal(&temp); err != nil {
		return err
	}
	for interfaceName, _ := range temp {
		switch interfaceName {
		case "FileEngine":
			var temp map[string]FileEngine
			if err := unmarshal(&temp); err != nil {
				return err
			}

		}
	}
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
	Artifact Artifact          `json:"artifact",yaml:"artifact"`
	Inputs   []Input           `json:"inputs",yaml:"inputs"`
	Outputs  map[string]Output `json:"inputs",yaml:"inputs"`
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

func (e *FileEngine) Do() chan State {
	c := make(chan State)
	return c
}

func (e *FileEngine) GetOutput() map[string]Output {
	return map[string]Output{}
}
