package choreography

import (
	"log"
)

// Node is a "runable" node description
type Node struct {
	ID       int               `json:"id"`
	State    int               `json:"state,omitempty"`
	Name     string            `json:"name,omitempty"`   // The targeted host
	Target   string            `json:"target,omitempty"` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Engine   string            `json:"engine,omitempty"` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Artifact string            `json:"artifact"`
	Inputs   []string          `json:"args,omitempty"`   // the arguments of the artifact, if needed
	Outputs  map[string]string `json:"output,omitempty"` // the key is the name of the parameter, the value its value (always a string)
	GraphID  string            `json:"-"`
}

// publishState in the etcd cluster
func (n *Node) publishState() {

}

func (n *Node) Monitor(event ...int) chan map[int]int {
	evt := make(chan map[int]int)
	return evt
}

// The output chan is filled with the desired state of the node
// In this function we ilplement the "intelligency" of the node,
// It where is decides if it needs to change its state or not
func (n *Node) RequiredState() chan int {
	state := make(chan int)

	// Find what event to monitor by looking in the matrix m

	// e is an array of all the ids to monitor
	var e []int
	var newEvt chan map[int]int

	// monitor the events in a seperate goroutine
	newEvt = n.Monitor(e...)
	log.Println("%v", newEvt)

	go func() {
		for newEvt := range newEvt {
			run := true
			for _, val := range newEvt {
				log.Println(val)
			}
			if run {
				state <- Running
			}
		}
	}()
	return state
}
