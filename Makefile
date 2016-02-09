TARGET:=./dist/$(shell uname -s)/$(shell uname -p)


$(TARGET)/choreography/choreography: choreography/*.go choreography/engines/*.go main.go
	go build -o $(TARGET)/choreography/choreography main.go

dist:\
	$(TARGET)/choreography/choreography\

clean:
	rm -rf $(TARGET)

testing: dist
