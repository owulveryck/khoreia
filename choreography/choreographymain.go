package choreography

import (
	"encoding/json"
	"log"
	"os"
)

func Main() {
	// Reading from stdin if opened
	var g Graph
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&g); err != nil {
		log.Panic(err)
	}
	log.Println(g)

	// create a cluster composed of all "targets"

}
