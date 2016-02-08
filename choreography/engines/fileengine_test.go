package engines

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	testFile := "/tmp/bla"
	f := &FileEngine{File: testFile}
	done := make(chan struct{})
	check := f.Check(done)
	go func() {
		for evt := range check {
			t.Log("Event: ", evt)
		}
	}()
	for i := 0; i < 10; i++ {
		// create the file /tmp/bla
		t.Log("Creating file")
		d1 := []byte("hello go\n")
		err := ioutil.WriteFile(testFile, d1, 0644)
		if err != nil {
			t.Log(err)
		}
		// remove the file /tmp/bla
		t.Log("Removing file")
		os.Remove(testFile)
	}

	// stop the routine
	time.Sleep(100 * time.Millisecond)

	done <- struct{}{}
}
