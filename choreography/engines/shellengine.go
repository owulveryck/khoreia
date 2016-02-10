package engines

import (
	"bytes"
	"fmt"
	"github.com/owulveryck/khoreia/choreography/event"
	"golang.org/x/net/context"
	"log"
	"os/exec"
)

type ShellEngine struct {
	Artifact string
	Args     []string
	outputs  map[string]string
}

func NewShellEngine(i map[string]interface{}) (*ShellEngine, error) {
	var artifact string
	for k, v := range i {
		switch k {
		case "artifact":
			artifact = v.(string)
		default:
			return nil, fmt.Errorf("Unknown field %v (%v) for  engine ShellEngine", k, v)
		}
	}
	return &ShellEngine{Artifact: artifact}, nil
}

func (s *ShellEngine) Do(ctx context.Context) {
	d := exec.Command(s.Artifact, s.Args...)
	// Set the stdin stdout and stderr of the dot subprocess
	stdinOfDotProcess, err := d.StdinPipe()
	if err != nil {
		return
	}
	defer stdinOfDotProcess.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line
	readCloser, err := d.StdoutPipe()
	if err != nil {
		return

	}

	// Actually run the dot subprocess
	if err = d.Run(); err != nil { //Use start, not run
		log.Println("An error occured: ", err) //replace with logger, or anything you want
		return
	}
	//fmt.Fprintf(stdinOfDotProcess, s)
	stdinOfDotProcess.Close()

	// Read from stdout and store it in the correct structure
	var buf bytes.Buffer
	buf.ReadFrom(readCloser)

	return
}

func (f *ShellEngine) Check(ctx context.Context, stop chan struct{}) chan *event.Event {
	c := make(chan *event.Event)
	return c
}

func (e *ShellEngine) GetOutput(ctx context.Context) interface{} {
	return nil
}
