GOPATH=$(shell go env GOPATH)

all: kubetool

kubetool:
	go build -o kubetool -buildvcs=false
	chmod ugo+x kubetool
