package choreography

import (
	"log"
	"time"

	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/etcdmain"
	"net/http"
)

var cfg = client.Config{
	Endpoints: []string{"http://127.0.0.1:2379"},
	Transport: client.DefaultTransport,
	// set timeout per request to fail fast when the target endpoint is unavailable
	HeaderTimeoutPerRequest: time.Second,
}

var kapi client.KeysAPI

func startEtcd() {
	go etcdmain.Main()
	ping := fmt.Sprintf("%v/v2/stats/leader", cfg.Endpoints[0])
	r, err := http.Get(ping)
	for err != nil || r.StatusCode != 200 {
		r, err = http.Get(ping)
		time.Sleep(100 * time.Millisecond)
	}
	log.Println("etcd up and running")

}

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
