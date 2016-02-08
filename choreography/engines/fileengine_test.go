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
			if !evt {
				t.Log("Creating file")
				f.Do()
			}
		}
	}()
	// Adding file
	err := ioutil.WriteFile(f.File, []byte{}, 0644)
	if err != nil {
		t.Log(err)
	}
	for i := 0; i < 10; i++ {
		// remove the file /tmp/bla
		t.Log("Removing file")
		os.Remove(testFile)
		// stop the routine
	}

	time.Sleep(1000 * time.Millisecond)
	done <- struct{}{}
}
