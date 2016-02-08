package engines

import (
	"fmt"
	"golang.org/x/exp/inotify"
	"io/ioutil"
	"log"
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
func (f *FileEngine) Check(stop chan struct{}) chan bool {
	c := make(chan bool)
	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)

	}

	log.Println("Watching", f.File)
	err = watcher.Watch(filepath.Dir(f.File))
	if err != nil {
		log.Fatal(err)

	}
	go func() {
		defer close(c)
		for {
			select {
			case <-stop:
				log.Println("Stop")
				return
			case ev := <-watcher.Event:
				if ev.Name == f.File {
					switch ev.Mask {
					case inotify.IN_CREATE:
						c <- true
					case inotify.IN_DELETE:
						c <- false
					}
				}
			case err := <-watcher.Error:
				log.Println("error:", err)

			}
		}
	}()
	return c
}
func (e *FileEngine) Do() {
	err := ioutil.WriteFile(e.File, []byte{}, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (e *FileEngine) GetOutput() interface{} {
	return nil
}
