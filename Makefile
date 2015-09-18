PACKAGES := \
	github.com/atulmirajkar/atgo/model \
	github.com/atulmirajkar/atgo/controller
DEPENDENCIES :=  

all: install

install: deps build
	go install model/model.go
	go install controller/controller.go 
build:
	#go get -d github.com/atulmirajkar/RPC-golang
	rm -rf $(GOPATH)/src/github.com/atgo
	mkdir -p $(GOPATH)/src/github.com/atulmirajkar/atgo
	cp -r ./* $(GOPATH)/src/github.com/atulmirajkar/atgo 

deps:
	go get $(DEPENDENCIES)

clean:
	go clean  $(PACKAGES) $(DEPENDENCIES)
	rm -rf $(GOBIN)/testclient
	rm -rf $(GOBIN)/testserver
	rm -rf $(GOPATH)/src/github.com/atulmirajkar/atgo
