This is a tool to generate two self-signed certificates for communication between choreography and executor

To generate the certificates for localhost:

```shell
go run generate_cert.go -host 127.0.0.1 -ca -target choreography
go run generate_cert.go -host 127.0.0.1 -ca -target executor
```


This will generate 4 files:

* choreography.pem
* choreography_key.pem
* executor.pem
* executor_key.pem

for mutual authentication, you need to use three files per service:

__ Orchestrator __
* choreography.pem
* choreography_key.pem
* executor.pem // To extract the root certificates

__ Executor __
* executor.pem
* executor_key.pem
* choreography.pem // To extract the root certificates

## See the graph_test for example
