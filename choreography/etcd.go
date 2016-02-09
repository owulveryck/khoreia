package choreography

import (
	"github.com/owulveryck/khoreia/choreography/event"
	"log"
	"time"

	etcd "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

var cfg = etcd.Config{
	Endpoints: []string{"http://127.0.0.1:2379"},
	Transport: etcd.DefaultTransport,
	// set timeout per request to fail fast when the target endpoint is unavailable
	HeaderTimeoutPerRequest: time.Second,
}

var kapi etcd.KeysAPI

func InitEtcd() {

	if kapi != nil {
		return
	}
	c, err := etcd.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi = etcd.NewKeysAPI(c)
}

func etcdWatch(ctx context.Context, w etcd.Watcher) chan *event.Event {
	c := make(chan *event.Event)
	go func() {
		for {
			resp, err := w.Next(ctx) // blocks here
			if err != nil {
				log.Println(err)
			}
			var ret bool
			switch resp.Node.Value {
			case "true":
				ret = true
			case "false":
				ret = false
			default:
				log.Printf("Unknown value %v, must be ai boolean", resp.Node.Value)
			}
			log.Println("Event received", resp)
			c <- &event.Event{ret, ""}
		}
	}()
	return c
}
