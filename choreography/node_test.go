package choreography

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestUnmarshalYAML(t *testing.T) {
	var dat []Node
	f, err := ioutil.ReadFile("test/topology.yaml")

	if err != nil {
		t.FailNow()

	}

	err = yaml.Unmarshal(f, &dat)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	//t.Log(dat)
}
