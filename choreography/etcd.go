package choreography

import (
	"log"
	"time"

	"github.com/coreos/etcd/client"
)

var cfg = client.Config{
	Endpoints: []string{"http://127.0.0.1:2379"},
	Transport: client.DefaultTransport,
	// set timeout per request to fail fast when the target endpoint is unavailable
	HeaderTimeoutPerRequest: time.Second,
}

var kapi client.KeysAPI

func initEtcdClient() {

	if kapi != nil {
		return
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi = client.NewKeysAPI(c)
}
