echo "Generating Orchestrator certificate"
go run generate_cert.go -ca -host 127.0.0.1 -target choreography
echo "Generating Executor certificate"
go run generate_cert.go -ca -host 127.0.0.1 -target executor
