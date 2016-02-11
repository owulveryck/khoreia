package engines

import (
	"fmt"
	"github.com/owulveryck/khoreia/choreography/event"
	"golang.org/x/net/context"
	"gopkg.in/fsnotify.v1"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/* Package engines implements the Interface mechanism*/

// FileEngine takes a single file as argument, it checks for its presence or create it
// It relies on the libnotify package
type FileEngine struct {
	File string `json:"path",yaml:"path"`
}

func NewFileEngine(i map[string]interface{}) (*FileEngine, error) {
	var file string
	for k, v := range i {
		switch k {
		case "path":
			file = v.(string)
		default:
			return nil, fmt.Errorf("Unknown field %v (%v) for  engine FileEngine", k, v)
		}
	}
	return &FileEngine{File: file}, nil
}

// Check if f.File is present and send an event on the channel if it
// appears or disappear
func (f *FileEngine) Check(ctx context.Context, stop chan struct{}) chan *event.Event {
	c := make(chan *event.Event)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)

	}

	//log.Println("Watching", f.File)
	err = watcher.Add(filepath.Dir(f.File))
	if err != nil {
		log.Fatal(err)

	}
	go func() {
		defer close(c)
		// Send initial event
		if _, err := os.Stat(f.File); os.IsNotExist(err) {
			// path/to/whatever does not exist
			c <- &event.Event{IsDone: false, Msg: fmt.Sprintf("Initial check, %v is not present", f.File)}
		} else {
			c <- &event.Event{IsDone: true, Msg: fmt.Sprintf("Initial check, %v is present", f.File)}
		}
		for {
			select {
			case <-stop:
				log.Println("Stop")
				return
			case ev := <-watcher.Events:
				if ev.Name == f.File {
					var evt *event.Event
					//log.Println("event:", ev)
					/*
						if ev.Op&fsnotify.Write == fsnotify.Write {
							log.Println("modified file:", ev.Name)
						}
					*/
					evt = &event.Event{IsDone: true, Msg: fmt.Sprintf("Event %v ", ev)}
					//log.Println("Event op:", ev.Op)
					switch ev.Op {
					case fsnotify.Create:
						evt.IsDone = true
						c <- evt
					case fsnotify.Remove:
						evt.IsDone = false
						c <- evt
					}
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)

			}
		}
	}()
	return c
}
func (e *FileEngine) Do(ctx context.Context) {
	//log.Println("Writing file", e)
	err := ioutil.WriteFile(e.File, []byte{}, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (e *FileEngine) GetOutput(ctx context.Context) interface{} {
	return nil
}
