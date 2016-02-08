package engines

import (
	"log"
)

/* Package engines implements the Interface mechanism*/

// FileEngine takes a single file as argument, it checks for its presence or create if
type FileEngine struct {
	File string `json:"file",yaml:"file"`
}

func NewFileEngine(i map[string]interface{}) *FileEngine {
	var file string
	for k, v := range i {
		switch k {
		case "file":
			file = v.(string)
		}
	}
	return &FileEngine{File: file}
}

// Check if f.File is present and send an event on the channel if it
// appears or disappear
func (f *FileEngine) Check() chan bool {
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
func (e *FileEngine) Do() {
	log.Println("Do Method...", e)
}

func (e *FileEngine) GetOutput() interface{} {
	return nil
}
