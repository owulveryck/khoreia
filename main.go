package main

import (
	"fmt"
	"github.com/owulveryck/khoreia/choreography"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	var nodes []choreography.Node
	f, err := ioutil.ReadFile("samples/topology.yaml")

	if err != nil {
		panic(err)

	}

	err = yaml.Unmarshal(f, &nodes)
	if err != nil {
		log.Println(err)
	}

	// Init the etcd client
	choreography.InitEtcd()
	// Generate a uuid
	u1 := uuid.NewV4()
	ctx := context.TODO()

	// Temp: for debug purpose
	for _, node := range nodes {
		for k, v := range node.Interfaces {
			etcdPath := fmt.Sprintf("/%s/nodes/%s/%s", u1, node.ID, k)
			//dependencies := []string{"/", "//"}
			v.Run(ctx, etcdPath, etcdPath)
		}
	}
	log.Println("Let's dance")

	stop := make(chan struct{})
	<-stop
	//time.Sleep(10 * time.Second)
}
