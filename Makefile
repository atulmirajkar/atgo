PACKAGES := \
	github.com/atulmirajkar/atgo/model \
	github.com/atulmirajkar/atgo/controller
DEPENDENCIES :=  

all: install

install: deps build release
	go install main.go
build:
	rm -rf $(GOPATH)/src/github.com/atgo
	mkdir -p $(GOPATH)/src/github.com/atulmirajkar/atgo
	cp -r ./* $(GOPATH)/src/github.com/atulmirajkar/atgo 
release:
	kill -9 $$(netstat -lpn | egrep -o [0-9]*/atgo | awk 'BEGIN { FS = "/" } ; { print $$1 }')
deps:
	go get $(DEPENDENCIES)

clean:
	go clean  $(PACKAGES) $(DEPENDENCIES)
	rm -rf $(GOBIN)/atgo
	rm -rf $(GOPATH)/src/github.com/atulmirajkar/atgo
