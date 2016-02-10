package main

import (
	"flag"
	"fmt"
	"github.com/owulveryck/khoreia/choreography"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	//var myself = flag.String("host", "local", "The hostname of this node")
	var topology = flag.String("topology", "samples/topology.yaml", "The topology file")
	flag.Parse()
	var nodes []choreography.Node
	f, err := ioutil.ReadFile(*topology)

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
			var dep []string

			etcdPath := fmt.Sprintf("/%s/nodes/%s/%s", u1, node.ID, k)
			// TODO, implement the lifecycle
			if k == "configure" {
				dep = append(dep, fmt.Sprintf("/%s/nodes/%s/%s", u1, node.ID, "create"))
			}
			for _, a := range node.Deps {
				for _, d := range a["nodes"] {
					dep = append(dep, fmt.Sprintf("/%s/nodes/%s/%s", u1, d, k))
				}
			}

			if dep == nil {
				v.Run(ctx, etcdPath)
			} else {
				v.Run(ctx, etcdPath, dep...)

			}
		}
	}
	log.Println("Let's dance")

	stop := make(chan struct{})
	<-stop
}
