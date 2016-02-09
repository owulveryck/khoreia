// Olivier Wulveryck - author of khoreia
// Copyright (C) 2015 Olivier Wulveryck
//
// This file is part of the khoreia project and
// is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package choreography

import (
	"fmt"
	etcd "github.com/coreos/etcd/client"
	"github.com/owulveryck/khoreia/choreography/engines"
	"github.com/owulveryck/khoreia/choreography/event"
	"golang.org/x/net/context"
	//"log"
	"reflect"
	"strconv"
)

// Objects implementing the Implementer interface will get their method called by a node
// when needed by its Interface
type Implementer interface {
	Do(context.Context) // Actuall
	Check(context.Context, chan struct{}) chan *event.Event
	GetOutput(context.Context) interface{}
}

type Interface struct {
	Do    Implementer `json:"do"`
	Check Implementer `json:"check"`
}

// Run calls Check.Check() (which runs in a goroutine) and wait fo all the conditions
// to be ok to call a Do.Do()
func (i *Interface) Run(ctx context.Context, etcdPath string, dependencies ...string) chan struct{} {
	stop := make(chan struct{})
	go func(i Interface) {
		done := make(chan struct{})

		chans := make([]chan *event.Event, len(dependencies))
		for i, dep := range dependencies {
			watcher := kapi.Watcher(dep, &etcd.WatcherOptions{Recursive: false})
			chans[i] = etcdWatch(ctx, watcher)
		}

		check := i.Check.Check(ctx, done)
		chans = append(chans, check)
		// Reference http://play.golang.org/p/8zwvSk4kjx
		cases := make([]reflect.SelectCase, len(chans))
		allCheck := make(map[chan *event.Event]bool, len(chans))
		for i, ch := range chans {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
			allCheck[ch] = false
		}

		remaining := len(cases)
		for remaining > 0 {
			chosen, value, ok := reflect.Select(cases)
			if !ok {
				// The chosen channel has been closed, so zero out the channel to disable the case
				cases[chosen].Chan = reflect.ValueOf(nil)
				remaining -= 1
				continue

			}
			//log.Printf("Read from channel %#v and received %s\n", chans[chosen], value)
			evt := value.Elem().FieldByName("IsDone").Bool()
			if chans[chosen] == check {

				// If its my own check, I can run if its true
				allCheck[chans[chosen]] = !evt
				kapi.Set(ctx, etcdPath, strconv.FormatBool(evt), nil)
			} else {
				allCheck[chans[chosen]] = evt
			}
			run := true
			for _, r := range allCheck {
				if r == false {
					run = false
				}
			}
			if run {
				i.Do.Do(ctx)
			}
		}

		i.Do.Do(ctx)
		for {
			select {
			case <-stop:
				return
			case evt := <-check:
				kapi.Set(ctx, etcdPath, strconv.FormatBool(evt.IsDone), nil)
				if !evt.IsDone {
					// TODO check if all the conditions are met
					i.Do.Do(ctx)
				}
			}
		}
	}(*i)
	return stop
}

// We need to specialised the Unmarshaling because of the Interfaces field
func (i *Interface) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// attributes is a key/value pair of attributes of an engine
	// eg: filePath = "/tmp/truc"
	type attributes map[string]interface{}
	// the intf is a map of map.
	// the first map's key is Do or Check
	// The second map's key is the engine type
	// The interface{} is all the definition of the engine to be passed to the
	// New function
	type intf map[string]map[string]attributes
	var t intf
	if err := unmarshal(&t); err != nil {
		return err
	}

	// This function takes a map as input
	// The key is the engine name and the input represent the arguments
	f := func(i map[string]attributes) (Implementer, error) {
		var implementer Implementer
		// Get the value of the engine
		for engine, value := range i {
			// Create a new implementer based on the engine name
			switch engine {
			case "file":
				impl, err := engines.NewFileEngine(value)
				if err != nil {
					return nil, err
				}
				implementer = impl
			case "nil":
				impl, err := NewNilEngine(value)
				if err != nil {
					return nil, err
				}
				implementer = impl
			}
		}

		return implementer, nil
	}
	// key = check | do, otherwise error
	for key, value := range t {
		switch key {
		case "do":
			val, err := f(value)
			if err != nil {
				return err
			}
			i.Do = val
		case "check":
			val, err := f(value)
			if err != nil {
				return err
			}
			i.Check = val
		default:
			return fmt.Errorf("Unknown method %v", key)
		}

	}
	return nil
}
