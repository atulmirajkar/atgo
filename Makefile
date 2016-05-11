PACKAGES := \
	github.com/atulmirajkar/atgo/model \
	github.com/atulmirajkar/atgo/controller
DEPENDENCIES :=

all: install

rcmi: release clean move install

install:
	go build -gcflags "-N -l" $(GOPATH)/src/github.com/atulmirajkar/atgo/model/model.go
	go build -gcflags "-N -l" $(GOPATH)/src/github.com/atulmirajkar/atgo/controller/controller.go
	go install -gcflags "-N -l"
move:
		mkdir -p $(GOPATH)/src/github.com/atulmirajkar/atgo
		cp -r ./* $(GOPATH)/src/github.com/atulmirajkar/atgo

clean:
			go clean  $(PACKAGES) $(DEPENDENCIES)
			rm -rf $(GOBIN)/atgo
			rm -rf $(GOPATH)/src/github.com/atulmirajkar/atgo

release:
	kill -9 $$(netstat -lpn | egrep -o [0-9]*/atgo | awk 'BEGIN { FS = "/" } ; { print $$1 }')
deps:
	go get $(DEPENDENCIES)
