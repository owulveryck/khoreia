TARGET:=./dist/$(shell uname -s)/$(shell uname -p)

executor: $(TARGET)/executor/executor

choreography: $(TARGET)/choreography/choreography

clients: $(TARGET)/clients/web

clients: $(TARGET)/clients/tosca $(TARGET)/clients/web

$(TARGET)/generate_cert: security/util/generate_cert.go
	go build -o $(TARGET)/generate_cert security/util/generate_cert.go

$(TARGET)/clients/tosca: clients/tosca/*.go ../toscalib/*.go ../toscalib/toscaexec/*.go
	go build -o $(TARGET)/clients/tosca/tosca2gorch clients/tosca/*.go

$(TARGET)/clients/web: clients/web/*.go clients/web/htdocs/* clients/web/tmpl/*
	go build -o $(TARGET)/clients/web/webclient clients/web/*go
	cp -r clients/web/htdocs clients/web/tmpl $(TARGET)/clients/web/

$(TARGET)/choreography/choreography: choreography/*.go http/*.go config/*.go
	go build -o $(TARGET)/choreography/choreography main.go

$(TARGET)/executor/sshConfig_sample.yaml: executor/sshConfig.yaml
	cp executor/sshConfig.yaml $(TARGET)/executor/sshConfig_sample.yaml

$(TARGET)/executor/executor: executor/*.go executor/cmd/*.go config/*.go
	go build -o $(TARGET)/executor/executor executor/cmd/main.go

$(TARGET)/certs/choreography.pem: $(TARGET)/certs/choreography_key.pem

$(TARGET)/certs/choreography_key.pem: $(TARGET)/generate_cert
	mkdir -p $(TARGET)/certs && \
	cd $(TARGET)/certs && \
	../generate_cert -ca -host 127.0.0.1 -target choreography

$(TARGET)/certs/executor.pem: $(TARGET)/certs/executor_key.pem

$(TARGET)/certs/executor_key.pem: $(TARGET)/generate_cert
	mkdir -p $(TARGET)/certs && \
	cd $(TARGET)/certs && \
	../generate_cert -ca -host 127.0.0.1 -target executor 

certificates: $(TARGET)/certs/choreography_key.pem $(TARGET)/certs/executor_key.pem

install_certificates: certificates $(TARGET)/choreography/choreography $(TARGET)/executor/executor
	cp $(TARGET)/certs/choreography*pem $(TARGET)/certs/executor.pem $(TARGET)/choreography && \
	cp $(TARGET)/certs/executor*pem $(TARGET)/certs/choreography.pem $(TARGET)/executor 
	
$(TARGET)/executor/config.json: config/executor.conf.sample
	cp config/executor.conf.sample $(TARGET)/executor/config.json

$(TARGET)/choreography/config.json: config/choreography.conf.sample
	cp config/choreography.conf.sample $(TARGET)/choreography/config.json

dist:\
	$(TARGET)/executor/executor\
	$(TARGET)/choreography/choreography\
	$(TARGET)/generate_cert\
	$(TARGET)/clients/web\
	install_certificates\
	clients\
	$(TARGET)/executor/sshConfig_sample.yaml\
	$(TARGET)/choreography/config.json\
	$(TARGET)/executor/config.json

clean:
	rm -rf $(TARGET)

testing: dist
	# Creating the layout
	tmux new-window -n "Gchoreography"
	tmux split-window -h 
	tmux split-window -v
	tmux select-pane -t 0
	tmux split-window -v
	tmux select-pane -t 0
	tmux resize-pane -D 10
	tmux send-keys "cd $(TARGET)/choreography && ./choreography" 
	tmux select-pane -t 3
	tmux send-keys "cd $(TARGET)/clients/web && ./webclient"
	tmux select-pane -t 2
	tmux send-keys "cd $(TARGET)/executor && ./executor"
	tmux select-pane -t 1
	tmux send-keys "$(TARGET)/clients/tosca/tosca2gorch -t clients/tosca/example/tosca_elk.yaml -i clients/tosca/example/inputs.yaml | curl  -X POST -H 'Content-Type:application/json' -H 'Accept:application/json' -d@- http://localhost:8080/v1/tasks" 
