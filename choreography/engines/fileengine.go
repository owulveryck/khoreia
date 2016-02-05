package engines

import (
	"log"
)

/* Package engines implements the Interface mechanism*/

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

func NewFileEngine(i map[string]interface{}) *FileEngine {
	var artifact Artifact
	for k, v := range i {
		switch k {
		case "artifact":
			artifact = Artifact(v.(string))
		}
	}
	return &FileEngine{artifact: artifact}
}

func (e *FileEngine) Do() chan State {
	c := make(chan State)
	log.Println("Do Method...", e)
	return c
}

func (e *FileEngine) GetOutput() map[string]Output {
	return map[string]Output{}
}
