package choreography

import (
	"github.com/owulveryck/khoreia/choreography/engines"
	"sync"
)

type Implementation interface {
	Do()
	Check() chan bool
	GetOutput() map[string]engines.Output
	GetWG() sync.WaitGroup // The wait  group will be used to tell check to wait while Do is working
}
