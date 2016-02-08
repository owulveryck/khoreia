package choreography

import (
	"github.com/owulveryck/khoreia/choreography/engines"
	//"sync"
)

// Objects implementing the Implementer interface will get their method called by a node
// when needed by its Interface
type Implementer interface {
	Do() // Actuall
	Check() chan bool
	GetOutput() map[string]engines.Output
	//GetWG() sync.WaitGroup // The wait  group will be used to tell check to wait while Do is working
}
